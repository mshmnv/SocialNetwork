package postgres

import (
	"database/sql"
	"os"

	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
)

const (
	conn = "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@db_container:$POSTGRES_PORT/social_network?sslmode=disable"
	//connAsync = "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@db_async_container:5432/social_network?sslmode=disable"
)

type DB struct {
	master db
	//async  DB
}

type db struct {
	db     *sql.DB
	closer func() error
}

func Connect() (*DB, error) {
	var err error
	dbSet := &DB{}

	dbSet.master.db, dbSet.master.closer, err = connect(conn)
	if err != nil {
		return nil, err
	}

	//dbSet.async.db, dbSet.async.closer, err = connect(connAsync)
	//if err != nil {
	//	return nil, nil, err
	//}

	logger.Info("Successfully connected to postgres")

	return dbSet, nil
}

func connect(connection string) (*sql.DB, func() error, error) {
	db, err := sql.Open("postgres", os.ExpandEnv(connection))
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Unable to connect to database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, err
	}

	return db, db.Close, nil
}

func (d *DB) GetConnection() *sql.DB {
	return d.master.db
}
