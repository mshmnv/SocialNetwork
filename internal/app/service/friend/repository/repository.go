package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	"github.com/pkg/errors"
)

type Repository struct {
	db *postgres.DB
}

const (
	friendTable    = "friends"
	pendingStatus  = "pending"
	approvedStatus = "approved"
)

func NewRepository(db *postgres.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SendFriendRequest(friendID uint64, userID uint64) error {
	tx, err := r.db.GetConnection().Begin()
	if err != nil {
		return err
	}

	userQuery, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(friendTable).
		SetMap(map[string]interface{}{
			"user_id":   userID,
			"friend_id": friendID,
			"status":    pendingStatus,
		}).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(userQuery, args...)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "Error adding new friend")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "Error adding new friend")
	}
	return nil
}

func (r *Repository) ApproveFriendRequest(friendID uint64, userID uint64) error {
	tx, err := r.db.GetConnection().Begin()
	if err != nil {
		return err
	}

	userQuery, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(friendTable).
		Set("status", approvedStatus).
		Where(sq.Eq{"user_id": userID, "friend_id": friendID}).ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(userQuery, args...)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "Error approving friend request")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "Error approving friend request")
	}

	return nil
}

func (r *Repository) DeleteFriend(friendID uint64, userID uint64) error {
	tx, err := r.db.GetConnection().Begin()
	if err != nil {
		return err
	}
	userQuery, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete(friendTable).
		Where(sq.Eq{"user_id": userID, "friend_id": friendID}).ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(userQuery, args...)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "Error deleting friend")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "Error deleting friend")
	}

	return nil
}

func (r *Repository) GetUserFriends(userID uint64) ([]uint64, error) {
	query1, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`user_id`).
		From(friendTable).
		Where(sq.Eq{"friend_id": userID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var result []uint64

	rows1, err := r.db.GetConnection().Query(query1, args...)
	if err != nil {
		return nil, errors.Wrap(err, "error getting user friends")
	}
	for rows1.Next() {
		var friendID uint64
		if err = rows1.Scan(&friendID); err != nil {
			return nil, errors.Wrap(err, "Error getting user friends")
		}
		result = append(result, friendID)
	}

	query2, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`friend_id`).
		From(friendTable).
		Where(sq.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows2, err := r.db.GetConnection().Query(query2, args...)
	if err != nil {
		return nil, errors.Wrap(err, "error getting user friends")
	}
	for rows2.Next() {
		var friendID uint64
		if err = rows2.Scan(&friendID); err != nil {
			return nil, errors.Wrap(err, "Error getting user friends")
		}
		result = append(result, friendID)
	}

	return result, nil
}
