package postgres

type User struct {
	ID                      int    `db:"id"`
	FirstName               string `db:"first_name"`
	LastName                string `db:"last_name,omitempty"`
	Username                string `db:"username,omitempty"`
	LanguageCode            string `db:"language_code,omitempty"`
	CanJoinGroups           bool   `db:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `db:"can_read_all_group_messages,omitempty"`
}

type Chat struct {
	ID            int    `db:"id"`
	Type          string `db:"type"`
	Title         string `db:"title"`
	Username      string `db:"username,omitempty"`
	FirstName     string `db:"first_name"`
	LastName      string `db:"last_name,omitempty"`
	SlowModeDelay int    `db:"slow_mode_delay,omitempty"`
}

type Zhuk struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Photo  string `db:"photo,omitempty"`
	UserID int    `db:"user_id"`
	ChatID int    `db:"chat_id"`
}
