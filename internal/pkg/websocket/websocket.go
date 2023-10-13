package websocket

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mshmnv/SocialNetwork/internal/pkg/clients"
	logger "github.com/sirupsen/logrus"
)

var webSocketMethod = "/post/feed/posted"

const pingPeriod = 5 * time.Second

var upgrader = websocket.Upgrader{}

func HTTPHandler() http.Handler {
	return Handler{}
}

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint64)
	if userID == 0 {
		logger.Errorf("Invalid user id")
	}

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Error upgrading", http.StatusInternalServerError)
		return
	}
	delivery := clients.AddFeedConnection(userID, connection)
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		clients.CloseFeedConnection(userID)
		ticker.Stop()
	}()

	for {
		select {
		case data, ok := <-delivery:
			if !ok {
				err := connection.WriteMessage(websocket.CloseMessage, []byte{})
				logger.Errorf("Unexpected closing connection: %s", err)
				return
			}

			err := connection.WriteMessage(websocket.TextMessage, data.Body)
			if err != nil {
				logger.Errorf("Error writing message to socket connection: %s", err)
				return
			}
		case <-ticker.C:
			if err := connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Errorf("Error ping: %s", err)
				return
			}
		}
	}
}
