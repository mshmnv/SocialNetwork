package rabbitmq

import (
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
)

const (
	url             = "amqp://guest:guest@rabbitmq/"
	queueNamePrefix = "service-queue"
	consumerName    = "service-consumer"
)

func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to connect to rabbit: %s", err)
	}
	logger.Info("Successfully connected to rabbit")

	return conn, nil
}
