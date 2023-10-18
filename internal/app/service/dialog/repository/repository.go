package repository

import (
	"database/sql"
	"sync"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/mshmnv/SocialNetwork/internal/app/service/dialog/datastruct"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type Repository struct {
	shardedDB *postgres.ShardedDB
	db        *postgres.DB
}

func NewRepository(sharded *postgres.ShardedDB, db *postgres.DB) *Repository {
	return &Repository{
		shardedDB: sharded,
		db:        db}
}

const (
	dialogTable    = "dialog"
	messagingTable = "messaging" // in sharded db
)

func (r *Repository) GetDialogID(user1, user2 int64) (int64, error) {
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`id`).
		From(dialogTable).
		Where(sq.Or{sq.Eq{"sender_id": user1, "receiver_id": user2}, sq.Eq{"sender_id": user2, "receiver_id": user1}}).
		ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "error getting dialog id")
	}

	var dialogID int64
	err = r.db.GetConnection().QueryRow(query, args...).Scan(&dialogID)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return r.CreateDialog(user1, user2)
		}
		return 0, errors.Wrap(err, "error getting dialog id")
	}
	return dialogID, nil
}

func (r *Repository) CreateDialog(sender, receiver int64) (int64, error) {
	tx, err := r.db.GetConnection().Begin()
	if err != nil {
		return 0, err
	}

	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(dialogTable).
		SetMap(map[string]any{
			"sender_id":   sender,
			"receiver_id": receiver}).
		Suffix("RETURNING id").ToSql()
	if err != nil {
		return 0, err
	}

	var dialogID int64
	err = tx.QueryRow(query, args...).Scan(&dialogID)
	if err != nil {
		_ = tx.Rollback()
		return 0, errors.Wrap(err, "Error adding dialog")
	}
	if err = tx.Commit(); err != nil {
		return 0, errors.Wrap(err, "Error adding dialog")
	}
	return dialogID, nil
}

func (r *Repository) Add(dialogID, sender, receiver int64, text string) error {
	c, err := r.shardedDB.GetConnection(dialogID)
	if err != nil {
		return err
	}

	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(messagingTable).
		SetMap(map[string]any{
			"dialog_id":   dialogID,
			"sender_id":   sender,
			"receiver_id": receiver,
			"text":        text,
			"sent_at":     time.Now()}).ToSql()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, args...)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "Error adding dialog message")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "Error adding dialog message")
	}
	return nil
}

func (r *Repository) GetDialogs(user int64) ([]int64, error) {
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`id`).
		From(dialogTable).
		Where(sq.Or{sq.Eq{"sender_id": user}, sq.Eq{"receiver_id": user}}).
		ToSql()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting dialog ids. Query [%s]", query)
	}

	var dialogIDs []int64
	rows, err := r.db.GetConnection().Query(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "error getting dialog ids")
	}
	for rows.Next() {
		var id int64
		if err = rows.Scan(&id); err != nil {
			return nil, errors.Wrap(err, "error getting dialog ids")
		}
		dialogIDs = append(dialogIDs, id)
	}

	return dialogIDs, nil
}

func (r *Repository) GetMessages(dialogIDs []int64) ([]datastruct.Message, error) {
	messages := make(chan []datastruct.Message, len(dialogIDs))
	errs := make(chan error, len(dialogIDs))

	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("sender_id", "receiver_id", "text", "sent_at").
		From(messagingTable)
	wg := sync.WaitGroup{}
	for _, dialog := range dialogIDs {
		wg.Add(1)
		go func(dialog int64) {
			defer wg.Done()
			r.getMessage(query, dialog, messages, errs)
		}(dialog)
	}

	results := make([]datastruct.Message, 0, 100)

	wg.Wait()
	close(messages)
	close(errs)

	for msg := range messages {
		results = append(results, msg...)
	}

	var err error
	for e := range errs {
		err = multierr.Append(err, e)
	}

	return results, err
}

func (r *Repository) getMessage(query sq.SelectBuilder, dialogID int64, messages chan []datastruct.Message, errs chan error) {
	c, err := r.shardedDB.GetConnection(dialogID)
	if err != nil {
		errs <- err
		return
	}
	q, args, err := query.Where(sq.Eq{"dialog_id": dialogID}).ToSql()
	if err != nil {
		errs <- err
		return
	}
	rows, err := c.Query(q, args...)
	if err != nil {
		errs <- err
		return
	}

	msgs := make([]datastruct.Message, 0, 1000)
	for rows.Next() {
		msg := datastruct.Message{DialogID: dialogID}
		if err = rows.Scan(&msg.Sender, &msg.Receiver, &msg.Text, &msg.SentAt); err != nil {
			errs <- err
			return
		}
		msgs = append(msgs, msg)
	}

	messages <- msgs
}
