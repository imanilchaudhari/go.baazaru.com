package models

// Auth table contains the information for each Auth
type Auth struct {
	Id        uint32 `db:"id"`
	User_id   string `db:"user_id"`
	Source    string `db:"source"`
	Source_id string `db:"source_id"`
}
