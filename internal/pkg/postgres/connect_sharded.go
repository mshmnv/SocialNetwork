package postgres

import (
	"database/sql"

	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
)

const shardCount = 2

const (
	Shard1 ShardNum = iota
	Shard2
)

type ShardNum int
type shards map[ShardNum]*sql.DB

type ShardedDB struct {
	db shards
}

var dsns = map[ShardNum]string{
	Shard1: "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@db_shard_1_container:$POSTGRES_PORT/social_network?sslmode=disable",
	Shard2: "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@db_shard_2_container:$POSTGRES_PORT/social_network?sslmode=disable",
}

func ConnectSharded() (*ShardedDB, error) {
	db := make(shards)

	for shardNum, dsn := range dsns {
		c, _, err := connect(dsn)
		if err != nil {
			logger.Fatalf("Error connection to db shard %d: %s", shardNum, err)
		}
		db[shardNum] = c
	}

	logger.Info("Successfully connected to sharded postgres")
	return &ShardedDB{db: db}, nil
}

func (s *ShardedDB) GetConnection(key int64) (*sql.DB, error) {
	if _, ok := s.db[ShardNum(key%shardCount)]; !ok {
		return nil, errors.New("Shard does not exists")
	}
	return s.db[ShardNum(key%shardCount)], nil
}
