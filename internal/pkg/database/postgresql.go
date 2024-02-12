package database

import (
	"database/sql"
	"fmt"

	"github.com/natanaelrusli/library-api-gin/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitPostgres(cfg *config.Config) (*sql.DB, error) {
	dbConfig := cfg.DBConfig

	dsn := fmt.Sprintf("host=%s user=%s password='' dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.DbName,
		dbConfig.Port,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
