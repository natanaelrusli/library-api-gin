package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_initDbConfig(t *testing.T) {
	tests := []struct {
		name  string
		want  DatabaseConfig
		mock  func()
		reset func()
	}{
		{
			name: "success",
			want: DatabaseConfig{
				Host:     "localhost",
				Port:     5432,
				DbName:   "library",
				Username: "postgres",
				Password: "",
			},
			mock: func() {
				os.Setenv("DATABASE_HOST", "localhost")
				os.Setenv("DATABASE_USERNAME", "postgres")
				os.Setenv("DATABASE_PASSWORD", "")
				os.Setenv("DATABASE_PORT", "5432")
				os.Setenv("DATABASE_NAME", "library")
			},
			reset: func() {
				os.Unsetenv("DATABASE_HOST")
				os.Unsetenv("DATABASE_USERNAME")
				os.Unsetenv("DATABASE_PASSWORD")
				os.Unsetenv("DATABASE_PORT")
				os.Unsetenv("DATABASE_NAME")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if got := initDbConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitConfig() = %v, want %v", got, tt.want)
			}
			tt.reset()
		})
	}
}

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name  string
		want  *Config
		mock  func()
		reset func()
	}{
		{
			name: "success",
			want: &Config{
				DBConfig: DatabaseConfig{
					Host:     "localhost",
					Port:     5432,
					DbName:   "library",
					Username: "postgres",
					Password: "",
				},
			},
			mock: func() {
				os.Setenv("DATABASE_HOST", "localhost")
				os.Setenv("DATABASE_USERNAME", "postgres")
				os.Setenv("DATABASE_PASSWORD", "")
				os.Setenv("DATABASE_PORT", "5432")
				os.Setenv("DATABASE_NAME", "library")
			},
			reset: func() {
				os.Unsetenv("DATABASE_HOST")
				os.Unsetenv("DATABASE_USERNAME")
				os.Unsetenv("DATABASE_PASSWORD")
				os.Unsetenv("DATABASE_PORT")
				os.Unsetenv("DATABASE_NAME")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exist := true
			_, err := os.Stat(".env")
			if os.IsNotExist(err) {
				os.Create(".env")
				exist = false
			}
			tt.mock()
			if got, err := InitConfig(); !reflect.DeepEqual(got, tt.want) {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}

			if !exist {
				os.Remove(".env")
			}
		})
	}
}

func TestInitConfigFail(t *testing.T) {
	unreadableEnvFilePath := ".env_unreadable"
	_, err := os.Stat(unreadableEnvFilePath)
	if os.IsNotExist(err) {
		os.Create(unreadableEnvFilePath)
	}
	defer os.Remove(unreadableEnvFilePath)

	os.Setenv("ENV_FILE_PATH", unreadableEnvFilePath)
	defer os.Unsetenv("ENV_FILE_PATH")

	got, err := InitConfig()
	assert.Error(t, err)
	assert.Nil(t, got)
}
