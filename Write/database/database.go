package database

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Database interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Close(ctx context.Context) error
}

type Storage struct {
	db *pgx.Conn
}

func ConfigStorage(url string) (*Storage, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		slog.Error("Error while connecting db", "error", err)
		return nil, err
	}
	return &Storage{db: conn}, nil
}

func (s *Storage) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return s.db.Query(ctx, sql, args...)
}
func (s *Storage) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return s.db.Exec(ctx, sql, args...)
}
func (s *Storage) Close(ctx context.Context) error {
	return s.db.Close(ctx)
}
