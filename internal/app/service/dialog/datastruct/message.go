package datastruct

import (
	"time"
)

type Message struct {
	DialogID int64     `db:"dialog_id"`
	Sender   uint64    `db:"sender_id"`
	Receiver uint64    `db:"receiver_id"`
	Text     string    `db:"text"`
	SentAt   time.Time `db:"sent_at"`
}
