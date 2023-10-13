package clients

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/mshmnv/SocialNetwork/internal/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
)

type IConsumer interface {
	GetDeliveryChannel() <-chan amqp.Delivery
	Close()
}

var c = &Clients{
	conns: make(map[uint64]connection, 0),
}

type Clients struct {
	mu    sync.RWMutex
	conns map[uint64]connection
}

type connection struct {
	conn     *websocket.Conn
	consumer IConsumer
}

func IsConnectionForUser(user uint64) bool {
	_, ok := c.conns[user]
	return ok
}

func AddFeedConnection(user uint64, conn *websocket.Conn) <-chan amqp.Delivery {
	consumer, err := rabbitmq.StartConsumerForConnection(user)
	if err != nil {
		logger.Errorf("Error starting consumer for user %d queue", user)
		return nil
	}
	c.mu.Lock()
	c.conns[user] = connection{conn: conn, consumer: consumer}
	c.mu.Unlock()

	return consumer.GetDeliveryChannel()
}

func CloseFeedConnection(user uint64) {
	if _, ok := c.conns[user]; ok {
		c.conns[user].consumer.Close()
		c.conns[user].conn.Close()
	}
	delete(c.conns, user)
}
