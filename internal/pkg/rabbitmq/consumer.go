package rabbitmq

import (
	"fmt"

	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
)

type Consumer struct {
	conn       *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
	deliveries <-chan amqp.Delivery
}

func (c *Consumer) Close() {
	if err := c.channel.Close(); err != nil {
		logger.Errorf("Error closing consumer channel: %v", err)
	}
	if err := c.conn.Close(); err != nil {
		logger.Errorf("Error closing consumer connection: %v", err)
	}
	logger.Infof("Consumer successfully stopped")
}

func (c *Consumer) GetDeliveryChannel() <-chan amqp.Delivery {
	return c.deliveries
}

func StartConsumerForConnection(userID uint64) (*Consumer, error) {
	c := &Consumer{}
	var err error
	c.conn, err = Connect()
	if err != nil {
		return nil, errors.Wrapf(err, "Error starting consumer. Dial")
	}
	go func() {
		logger.Infof("Closing consumer connection: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, errors.Wrapf(err, "Error starting consumer. Channel")
	}
	c.queue, err = c.channel.QueueDeclare(
		fmt.Sprintf("%s-%d", queueNamePrefix, userID),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "Error starting connection. Queue Declare")
	}
	c.deliveries, err = c.channel.Consume(
		c.queue.Name,
		fmt.Sprintf("%s-%d", consumerName, userID),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "Error starting connection. Queue Consume")
	}

	logger.Infof("Consumer successfully started")
	return c, nil
}
