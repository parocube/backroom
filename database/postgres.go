package database

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/parocube/backroom/config"
)

func NewPostgres(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", cfg.Postgres.PostgresURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Postgres.PostgresMaxOpenConn)
	db.SetMaxIdleConns(cfg.Postgres.PostgresMaxIdleConn)
	db.SetConnMaxLifetime(cfg.Postgres.PostgresConnMaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
