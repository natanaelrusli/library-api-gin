package postgres

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBorrowedRecord(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}

	repo := &postgresBorrowingRecordRepository{
		Conn: db,
	}

	columns := []string{
		"id",
		"user_id",
		"book_id",
		"status",
		"borrowing_date",
		"returning_date",
		"created_at",
		"updated_at",
		"deleted_at",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		1,
		1,
		22,
		"BORROWED",
		time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		nil,
		time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		nil,
	)

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM borrowing_records WHERE status = 'BORROWED' AND deleted_at IS NULL;")).
		WillReturnRows(rows)

	res, err := repo.GetAllBorrowedRecord(ctx)
	if err != nil {
		t.Fatalf("%v", err)
	}

	expectedResult := []domain.BorrowingRecord{
		{
			Id:            1,
			UserId:        1,
			BookId:        22,
			Status:        "BORROWED",
			BorrowingDate: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			ReturningDate: sql.NullTime{},
			CreatedAt:     time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			UpdatedAt:     time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			DeletedAt:     sql.NullTime{},
		},
	}

	assert.Equal(t, res, expectedResult)
}

func TestCreateRecord(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}
	repo := &postgresBorrowingRecordRepository{
		Conn: db,
	}

	defer db.Close()

	columns := []string{
		"id",
		"user_id",
		"book_id",
		"status",
		"borrowing_date",
		"returning_date",
		"created_at",
		"updated_at",
		"deleted_at",
	}

	record := domain.BorrowingRecord{
		Id:            1,
		UserId:        1,
		BookId:        22,
		Status:        "BORROWED",
		BorrowingDate: time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		ReturningDate: sql.NullTime{},
		CreatedAt:     time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		UpdatedAt:     time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		DeletedAt:     sql.NullTime{},
	}

	rows := sqlmock.NewRows(columns).AddRow(
		1,
		1,
		22,
		"BORROWED",
		time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		nil,
		time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
		nil,
	)

	mock.ExpectQuery(
		regexp.QuoteMeta(`
			INSERT INTO borrowing_records
			(user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at)
			VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
			RETURNING
			id, user_id, book_id, status, borrowing_date, returning_date, created_at, updated_at, deleted_at;
		`)).
		WithArgs(
			1,
			22,
			"BORROWED",
			time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			nil,
			time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC),
			nil,
		).
		WillReturnRows(rows)

	res, err := repo.CreateRecord(ctx, record)
	if err != nil {
		t.Fatalf("%v", err)
	}

	assert.NotNil(t, res)
	assert.Equal(t, res, record)
}
