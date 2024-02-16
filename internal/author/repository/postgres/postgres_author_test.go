package postgres

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natanaelrusli/library-api-gin/internal/domain"
)

func TestGetByID(t *testing.T) {
	var ctx context.Context
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}

	defer db.Close()

	repo := &postgresAuthorRepository{Conn: db}

	expectedRows := sqlmock.
		NewRows([]string{"id", "name", "created_at", "updated_at", "deleted_at"}).
		AddRow(
			1, "John Doe", time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC), time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC), nil,
		)

	mock.
		ExpectQuery("SELECT \\* from authors WHERE deleted_at IS NULL AND id = \\$1;").
		WithArgs(1).
		WillReturnRows(expectedRows)

	result, err := repo.GetByID(ctx, 1)

	if err != nil {
		t.Fatalf("error calling GetByID: %v", err)
	}

	expectedResult := domain.Author{
		ID:        1,
		Name:      "John Doe",
		CreatedAt: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		UpdatedAt: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		DeletedAt: sql.NullTime{},
	}

	if result != expectedResult {
		t.Errorf("Expected result %+v, got %+v", expectedResult, result)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
}
