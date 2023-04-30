package postgres

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
)

var dbKey = "db"

const (
	conn      = "postgres://admin:root@db_container:5432/social_network?sslmode=disable"
	connAsync = "postgres://admin:root@db_async_container:5432/social_network?sslmode=disable"
	//connSync = "postgres://admin:root@db_sync_1_container:5433/social_network?sslmode=disable"
)

type DatabaseSet struct {
	master DB
	async  DB
	//syncDB *DB
}

type DB struct {
	db     *sql.DB
	closer func() error
}

func Connect(ctx context.Context) (*DatabaseSet, context.Context, error) {
	var err error
	dbSet := DatabaseSet{}

	dbSet.master.db, dbSet.master.closer, err = connect(conn)
	if err != nil {
		return nil, nil, err
	}

	dbSet.async.db, dbSet.async.closer, err = connect(connAsync)
	if err != nil {
		return nil, nil, err
	}

	logger.Info("Successfully connected to database")

	return &dbSet, newContext(ctx, &dbSet), nil
}

func connect(connection string) (*sql.DB, func() error, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Unable to connect to database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, err
	}

	return db, db.Close, nil
}

func newContext(ctx context.Context, dbSet *DatabaseSet) context.Context {
	ctx = context.WithValue(ctx, &dbKey, dbSet)
	return ctx
}

func FromContext(ctx context.Context) *DatabaseSet {
	dbStorage, ok := ctx.Value(&dbKey).(*DatabaseSet)
	if !ok {
		panic("Error getting connection from context")
	}

	return dbStorage
}

func GetDB(ctx context.Context) *sql.DB {
	return FromContext(ctx).master.db
}

func GetAsyncDB(ctx context.Context) *sql.DB {
	return FromContext(ctx).async.db
}
