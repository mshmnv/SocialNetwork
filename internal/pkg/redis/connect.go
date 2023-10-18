package redis

import (
	"os"

	"github.com/go-redis/redis"
	logger "github.com/sirupsen/logrus"
)

type Client struct {
	*redis.Client
}

func Connect() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	logger.Info("Successfully connected to redis")

	return &Client{client}, nil
}
