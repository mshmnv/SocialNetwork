package repository

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/mshmnv/SocialNetwork/internal/app/service/user/datastruct"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	"github.com/pkg/errors"
)

type Repository struct {
	db *postgres.DB
}

const (
	userTable = "users"
)

func NewRepository(db *postgres.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Register(data *datastruct.User) error {
	tx, err := r.db.GetConnection().Begin()
	if err != nil {
		return err
	}

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
	if err != nil {
		return nil, err
	}

	var results datastruct.User

	if err = r.db.GetConnection().QueryRow(query, args...).
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
	if err != nil {
		return nil, err
	}

	var results datastruct.LoginData

	if err = r.db.GetConnection().QueryRow(query, args...).
		Scan(&results.ID, &results.Password); err != nil {
		return nil, errors.Wrap(err, "no user with this email found")
	}

	return &results, nil
}

func (r *Repository) IsExistedUser(id uint64) bool {
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`1`).
		From(userTable).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return false
	}

	var result int64
	if err = r.db.GetConnection().QueryRow(query, args...).
		Scan(&result); err != nil {
		return false
	}

	return true
}

const searchLimit = 100

func (r *Repository) Search(firstName, secondName string) ([]datastruct.User, error) {
	query, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select(`first_name, second_name, age, birthdate, biography, city`).
		From(userTable).
		Where(sq.Like{"upper(first_name)": strings.ToUpper(firstName) + "%", "upper(second_name)": strings.ToUpper(secondName) + "%"}).
		OrderBy("id").
		Limit(searchLimit).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "Error searching users")
	}

	rows, err := r.db.GetConnection().Query(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Error searching users")
	}
	defer rows.Close()

	var results []datastruct.User
	for rows.Next() {
		user := datastruct.User{}
		if err = rows.Scan(&user.FirstName, &user.SecondName, &user.Age, &user.BirthDate, &user.Biography, &user.City); err != nil {
			return nil, errors.Wrap(err, "Error searching users")
		}
		results = append(results, user)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "Error searching users")
	}

	return results, nil
}

func (r *Repository) AddUsers() error {
	f, err := os.Open("info/testing/people.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	byteValue, _ := io.ReadAll(f)

	var users []datastruct.User
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return err
	}

	return r.addUsers(users)
}

func (r *Repository) addUsers(data []datastruct.User) error {
	tx, err := r.db.GetConnection().Begin()
	if err != nil {
		return err
	}

	for _, user := range data {
		userQuery, args, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
			Insert(userTable).
			SetMap(user.GetUsersDBRecord()).ToSql()
		if err != nil {
			return err
		}

		_, err = tx.Exec(userQuery, args...)
		if err != nil {
			_ = tx.Rollback()
			return errors.Wrap(err, "Error adding user data")
		}
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "Error adding user data")
	}
	return nil
}
