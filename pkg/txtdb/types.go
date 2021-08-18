package txtdb

import (
	"fmt"
	"strconv"
)

type Tabler interface {
	Key() (string, string)
	Format() string
	ToString() string
}

// Chat
type Chat struct {
	ID            int    `txtdb:"id"`
	Type          string `txtdb:"type"`
	Title         string `txtdb:"title"`
	Username      string `txtdb:"username"`
	FirstName     string `txtdb:"first_name"`
	LastName      string `txtdb:"last_name"`
	SlowModeDelay int    `txtdb:"slow_mode_delay"`
}

func (c Chat) Key() (string, string) {
	return "id", strconv.FormatInt(int64(c.ID), 10)
}

func (c Chat) Format() string {
	return "%d\t%s\t%s\t%s\t%s\t%s\t%d"
}

func (c Chat) ToString() string {
	return fmt.Sprintf(c.Format(), c.ID, c.Type, c.Title, c.Username, c.FirstName, c.LastName, c.SlowModeDelay)
}

// User
type User struct {
	ID                      int    `txtdb:"id"`
	Username                string `txtdb:"username"`
	FirstName               string `txtdb:"first_name"`
	LastName                string `txtdb:"last_name"`
	LanguageCode            string `txtdb:"language_code"`
	CanJoinGroups           bool   `txtdb:"can_join_groups"`
	CanReadAllGroupMessages bool   `txtdb:"can_read_all_group_messages"`
}

func (u User) Key() (string, string) {
	return "id", strconv.FormatInt(int64(u.ID), 10)
}

func (u User) Format() string {
	return "%d\t%s\t%s\t%s\t%s\t%t\t%t"
}

func (u User) ToString() string {
	return fmt.Sprintf(u.Format(), u.Username, u.FirstName, u.LastName, u.LanguageCode, u.CanJoinGroups, u.CanReadAllGroupMessages)
}

// Zhuk
type Zhuk struct {
	ID     int    `txtdb:"id"`
	Name   string `txtdb:"name"`
	Photo  string `txtdb:"photo"`
	UserID int    `txtdb:"user_id"`
	ChatID int    `txtdb:"chat_id"`
	RoleID int    `txtdb:"role_id"`
}

func (z Zhuk) Key() (string, string) {
	return "id", strconv.FormatInt(int64(z.ID), 10)
}

func (z Zhuk) Format() string {
	return "%d\t%s\t%s\t%d\t%d\t%d"
}

func (z Zhuk) ToString() string {
	return fmt.Sprintf(z.Format(), z.ID, z.Name, z.Photo, z.UserID, z.ChatID, z.RoleID)
}
