package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	"github.com/pkg/errors"
)

type Repository struct {
	ctx context.Context
}

const (
	friendTable    = "friends"
	pendingStatus  = "pending"
	approvedStatus = "approved"
)

func NewRepository(ctx context.Context) *Repository {
	return &Repository{ctx: ctx}
}

func (r *Repository) SendFriendRequest(friendID uint64, userID uint64) error {
	tx, err := postgres.GetDB(r.ctx).Begin()
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
	tx, err := postgres.GetDB(r.ctx).Begin()
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
	tx, err := postgres.GetDB(r.ctx).Begin()
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
