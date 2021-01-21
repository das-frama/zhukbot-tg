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
