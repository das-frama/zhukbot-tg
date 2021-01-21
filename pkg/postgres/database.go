package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Queryer interface {
	CreateUser(User) error
}

type db struct {
	pool *pgxpool.Pool
}

// New create/open a sqlite db file.
func New(pool *pgxpool.Pool) Queryer {
	return &db{
		pool: pool,
	}
}

func (db *db) CreateUser(u User) error {
	user := User{}
	err := db.pool.QueryRow(context.Background(), "SELECT * FROM users WHERE id=$1", u.ID).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.LanguageCode, &user.CanJoinGroups, &user.CanReadAllGroupMessages,
	)
	if err == pgx.ErrNoRows {
		tag, err := db.pool.Exec(
			context.Background(),
			"INSERT INTO users (id, first_name, last_name, username, language_code, can_join_groups, can_read_all_group_messages) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			u.ID, u.FirstName, u.LastName, u.Username, u.LanguageCode, u.CanJoinGroups, u.CanReadAllGroupMessages,
		)
		if err != nil {
			return err
		}

		if ok := tag.Insert(); !ok {
			return errors.New("error in inserting new user")
		}
	} else if err != nil {
		return errors.New("QueryRow failed: " + err.Error())
	}

	return nil
}
