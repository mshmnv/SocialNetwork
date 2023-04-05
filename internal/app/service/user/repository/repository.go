package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/mshmnv/SocialNetwork/internal/app/service/user/datastruct"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	"github.com/pkg/errors"
)

type Repository struct {
	ctx context.Context
}

const (
	userTable = "users"
)

func NewRepository(ctx context.Context) *Repository {
	return &Repository{ctx: ctx}
}

func (r *Repository) Register(data *datastruct.User) error {

	tx, err := postgres.GetDB(r.ctx).Begin()

	userQuery, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(userTable).
		SetMap(data.GetUsersDBRecord()).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(userQuery, args...)
	if err != nil {
		_ = tx.Rollback()
		return errors.Wrap(err, "Error adding user data")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "Error adding user data")
	}
	return nil
}

func (r *Repository) GetUser(id uint64) (*datastruct.User, error) {
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`first_name, second_name, age, birthdate, biography, city`).
		From(userTable).
		Where(sq.Eq{"id": id}).
		ToSql()

	var results datastruct.User

	if err = postgres.GetDB(r.ctx).QueryRow(query, args...).
		Scan(&results.FirstName, &results.SecondName, &results.Age, &results.BirthDate, &results.Biography, &results.City); err != nil {
		return nil, errors.Wrap(err, "Error getting user data")
	}

	return &results, nil
}

func (r *Repository) GetLoginData(id uint64) (*datastruct.LoginData, error) {
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`id, password`).
		From(userTable).
		Where(sq.Eq{"id": id}).
		ToSql()

	var results datastruct.LoginData

	if err = postgres.GetDB(r.ctx).QueryRow(query, args...).
		Scan(&results.ID, &results.Password); err != nil {
		return nil, errors.Wrap(err, "No user with this email found. Register!")
	}

	return &results, nil
}

func (r *Repository) IsExistedUser(id uint64) bool {
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`1`).
		From(userTable).
		Where(sq.Eq{"id": id}).
		ToSql()

	var result int64
	if err = postgres.GetDB(r.ctx).QueryRow(query, args...).
		Scan(&result); err != nil {
		return false
	}

	return true
}
