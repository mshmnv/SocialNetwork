package datastruct

import (
	"encoding/json"
	"time"
)

type Post struct {
	PostID    uint64    `db:"id" json:"-"`
	AuthorID  uint64    `db:"author_id" json:"author_id"`
	Text      string    `db:"text" json:"text"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	IsDeleted bool      `db:"is_deleted" json:"-"`
}

func (p Post) Encode() ([]byte, error) {
	res, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return res, nil
}
