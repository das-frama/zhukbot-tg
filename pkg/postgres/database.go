package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Queryer interface {
	CreateUser(User) error
	CreateChat(Chat) error
	CreateZhuk(Zhuk) (int, error)
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
	var ok bool
	err := db.pool.QueryRow(context.Background(), "SELECT 1 FROM users WHERE id=$1", u.ID).Scan(&ok)
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

func (db *db) CreateChat(c Chat) error {
	var ok bool
	err := db.pool.QueryRow(context.Background(), "SELECT 1 FROM chats WHERE id=$1", c.ID).Scan(&ok)
	if err == pgx.ErrNoRows {
		var id int
		err = db.pool.QueryRow(
			context.Background(),
			"INSERT INTO chats (id, type, title, username, first_name, last_name, slow_mode_delay) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
			c.ID, c.Type, c.Title, c.Username, c.FirstName, c.LastName, c.SlowModeDelay,
		).Scan(&id)
		if err != nil {
			return err
		}
	} else if err != nil {
		return errors.New("QueryRow failed: " + err.Error())
	}

	return nil
}

func (db *db) CreateZhuk(z Zhuk) (id int, err error) {
	err = db.pool.QueryRow(context.Background(), "SELECT id FROM zhuks WHERE user_id=$1 AND chat_id=$2", z.UserID, z.ChatID).Scan(&id)
	if err == pgx.ErrNoRows {
		err := db.pool.QueryRow(
			context.Background(),
			"INSERT INTO zhuks (user_id, chat_id, name) VALUES ($1, $2, $3) RETURNING id",
			z.UserID, z.ChatID, z.Name,
		).Scan(&id)
		if err != nil {
			return id, err
		}
	}

	return id, nil
}
