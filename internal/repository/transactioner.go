package repository

import (
	"context"
	"database/sql"
)

type SqlGdbc interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type Transactioner interface {
	Rollback() error
	Commit() error
	BeginTx() (Transactioner, error)
	BookRepository() BookRepository
}

type SqlTransaction struct {
	db *sql.DB
	tx *sql.Tx
}

func NewSqlTransaction(db *sql.DB) *SqlTransaction {
	return &SqlTransaction{
		db: db,
	}
}

func (s *SqlTransaction) BeginTx() (Transactioner, error) {
	tx, err := s.db.Begin()
	return &SqlTransaction{db: s.db, tx: tx}, err
}

func (s *SqlTransaction) Rollback() error {
	return s.tx.Rollback()
}

func (s *SqlTransaction) Commit() error {
	return s.tx.Commit()
}

func (s *SqlTransaction) BookRepository() BookRepository {
	return &postgresBookRepository{
		Conn: s.tx,
	}
}
