package repository

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/mshmnv/SocialNetwork/internal/app/service/post/datastruct"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	"github.com/pkg/errors"
)

type Repository struct {
	db *postgres.DB
}

const (
	postTable = "posts"
)

func NewRepository(db *postgres.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(userID uint64, text string) error {
	tx, err := r.db.GetConnection().Begin()
	if err != nil {
		return err
	}

	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(postTable).
		SetMap(map[string]any{
			"author_id":  userID,
			"text":       text,
			"created_at": time.Now(),
			"updated_at": time.Now(),
			"is_deleted": false}).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "Error adding post")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "Error adding post")
	}
	return nil
}

func (r *Repository) Update(userID, postID uint64, text string) error {
	tx, err := r.db.GetConnection().Begin()
	if err != nil {
		return err
	}

	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(postTable).
		SetMap(map[string]any{
			"text":       text,
			"updated_at": time.Now()}).
		Where(sq.Eq{
			"author_id": userID,
			"id":        postID}).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "Error updating post")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "Error updating post")
	}
	return nil
}
func (r *Repository) Delete(userID, postID uint64) error {
	tx, err := r.db.GetConnection().Begin()
	if err != nil {
		return err
	}
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(postTable).
		Set("is_deleted", true).
		Where(sq.Eq{
			"author_id": userID,
			"id":        postID}).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, args...)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "Error deleting post")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "Error deleting post")
	}

	return nil
}

func (r *Repository) Get(postID uint64) (*datastruct.Post, error) {
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`id, author_id, text`).
		From(postTable).
		Where(sq.Eq{"id": postID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var results datastruct.Post

	if err = r.db.GetConnection().QueryRow(query, args...).
		Scan(&results.PostID, &results.AuthorID, &results.Text); err != nil {
		return nil, errors.Wrap(err, "error getting post")
	}

	return &results, nil
}

func (r *Repository) GetPostsAfterDate(date time.Time) ([]datastruct.Post, error) {
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`id, author_id, text, created_at, updated_at, is_deleted`).
		From(postTable).
		Where(sq.Gt{"updated_at": date}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var result []datastruct.Post
	rows2, err := r.db.GetConnection().Query(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "error getting posts after date 1")
	}
	for rows2.Next() {
		post := datastruct.Post{}
		if err = rows2.Scan(&post.PostID, &post.AuthorID, &post.Text, &post.CreatedAt, &post.UpdatedAt, &post.IsDeleted); err != nil {
			return nil, errors.Wrap(err, "Error getting posts after date 2")
		}
		result = append(result, post)
	}

	return result, nil
}
