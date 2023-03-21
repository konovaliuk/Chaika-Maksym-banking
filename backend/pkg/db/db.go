package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/uptrace/bun/dialect/pgdialect"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
)

func NewCockroachDB(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Addr,
		cfg.Database,
	)

	sqldb, err := setupDB(cfg.DriverName, dsn, cfg)
	if err != nil {
		return nil, err
	}

	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	return sqldb, nil
}

func NewSingleTransactionDB(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Addr,
		cfg.Database,
	)

	txdb.Register("txdb", cfg.DriverName, dsn)

	sqldb, err := setupDB("txdb", dsn, cfg)
	if err != nil {
		return nil, err
	}

	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	return sqldb, nil
}

func NewBunDB(cfg Config) (*bun.DB, error) {
	sqldb, err := NewCockroachDB(cfg)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func setupDB(driverName string, dsn string, cfg Config) (*sql.DB, error) {
	sqldb, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	sqldb.SetMaxOpenConns(cfg.MaxOpenConnections)
	sqldb.SetMaxIdleConns(cfg.MaxIdleConnections)
	sqldb.SetConnMaxLifetime(time.Minute * time.Duration(cfg.MaxConnectionLifetimeMinutes))

	return sqldb, nil
}
