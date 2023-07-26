package redis

import (
	"context"

	"github.com/go-redis/redis"
	logger "github.com/sirupsen/logrus"
)

var redisKey = "redis"

const (
	host = "cache"
	port = "6379"
	pass = "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81"
)

func Connect(ctx context.Context) (context.Context, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	logger.Info("Successfully connected to database")

	return newContext(ctx, client), nil
}

func newContext(ctx context.Context, redis *redis.Client) context.Context {
	ctx = context.WithValue(ctx, &redisKey, redis)
	return ctx
}

func FromContext(ctx context.Context) *redis.Client {
	dbStorage, ok := ctx.Value(&redisKey).(*redis.Client)
	if !ok {
		panic("Error getting redis connection from context")
	}

	return dbStorage
}

func GetRedis(ctx context.Context) *redis.Client {
	return FromContext(ctx)
}
