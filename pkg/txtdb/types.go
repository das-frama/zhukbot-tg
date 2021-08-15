package txtdb

type Chat struct {
	ID            int    `txtdb:"id"`
	Type          string `txtdb:"type"`
	Title         string `txtdb:"title"`
	Username      string `txtdb:"username"`
	FirstName     string `txtdb:"first_name"`
	LastName      string `txtdb:"last_name"`
	SlowModeDelay int    `txtdb:"slow_mode_delay"`
}

type User struct {
	ID                      int    `txtdb:"id"`
	FirstName               string `txtdb:"first_name"`
	LastName                string `txtdb:"last_name"`
	Username                string `txtdb:"username"`
	LanguageCode            string `txtdb:"language_code"`
	CanJoinGroups           bool   `txtdb:"can_join_groups"`
	CanReadAllGroupMessages bool   `txtdb:"can_read_all_group_messages"`
}

type Zhuk struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Photo  string `db:"photo,omitempty"`
	UserID int    `db:"user_id"`
	ChatID int    `db:"chat_id"`
	RoleID int    `db:"chat_id,omitempty"`
}
