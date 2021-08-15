package txtdb

import "fmt"

type Stringer interface {
	ToString() string
}

type file struct {
	FileName string
}

type Chat struct {
	file
	ID            int    `txtdb:"id"`
	Type          string `txtdb:"type"`
	Title         string `txtdb:"title"`
	Username      string `txtdb:"username"`
	FirstName     string `txtdb:"first_name"`
	LastName      string `txtdb:"last_name"`
	SlowModeDelay int    `txtdb:"slow_mode_delay"`
}

func (c Chat) ToString() string {
	return fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%s\t%d", c.ID, c.Type, c.Title, c.Username, c.FileName, c.LastName, c.SlowModeDelay)
}

type User struct {
	file
	ID                      int    `txtdb:"id"`
	Username                string `txtdb:"username"`
	FirstName               string `txtdb:"first_name"`
	LastName                string `txtdb:"last_name"`
	LanguageCode            string `txtdb:"language_code"`
	CanJoinGroups           bool   `txtdb:"can_join_groups"`
	CanReadAllGroupMessages bool   `txtdb:"can_read_all_group_messages"`
}

func (u User) ToString() string {
	return fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%t\t%t", u.ID, u.Username, u.FirstName, u.LastName, u.LanguageCode, u.CanJoinGroups, u.CanReadAllGroupMessages)
}

type Zhuk struct {
	file
	ID     int    `txtdb:"id"`
	Name   string `txtdb:"name"`
	Photo  string `txtdb:"photo"`
	UserID int    `txtdb:"user_id"`
	ChatID int    `txtdb:"chat_id"`
	RoleID int    `txtdb:"role_id"`
}

func (z Zhuk) ToString() string {
	return fmt.Sprintf("%d\t%s\t%s\t%d\t%d\t%d", z.ID, z.Name, z.Photo, z.UserID, z.ChatID, z.RoleID)
}
