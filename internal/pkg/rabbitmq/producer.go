package rabbitmq

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
)

type Producer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func (p *Producer) Close() {
	if err := p.channel.Close(); err != nil {
		logger.Errorf("Error closing consumer channel: %v", err)
	}
	if err := p.conn.Close(); err != nil {
		logger.Errorf("Error closing consumer connection: %v", err)
	}
}

func NewProducer() (*Producer, error) {
	conn, err := Connect()
	if err != nil {
		return nil, errors.Wrapf(err, "Error starting producer. Dial")
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrapf(err, "Error starting producer. Channel")
	}
	logger.Infof("Producer is successfully started")
	return &Producer{conn: conn, channel: channel}, nil
}

func (p *Producer) Produce(message []byte, userID uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := p.channel.PublishWithContext(ctx,
		"",
		fmt.Sprintf("%s-%d", queueNamePrefix, userID),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return err
	}
	return nil
}
