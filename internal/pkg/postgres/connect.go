package postgres

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

var dbKey = "db"

const conn = "postgres://admin:root@db_container:5432/social_network?sslmode=disable"

// conn      "postgresql://test:test@localhost:5432/test")

func Connect() (*sql.DB, func() error, error) {

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Unable to connect to database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, err
	}

	return db, db.Close, nil
}

func NewContext(ctx context.Context, db *sql.DB) context.Context {
	dbCtx := context.WithValue(ctx, &dbKey, db)
	return dbCtx
}

func FromContext(ctx context.Context) *sql.DB {
	dbStorage, ok := ctx.Value(&dbKey).(*sql.DB)
	if !ok {
		panic("Error getting connection from context")
	}

	return dbStorage
}

func GetDB(ctx context.Context) *sql.DB {
	return FromContext(ctx)
}
