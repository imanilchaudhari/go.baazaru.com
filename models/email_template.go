package models

import "time"

// EmailTemplate table contains the information for each EmailTemplate
type EmailTemplate struct {
	Id          uint32    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Created_at  time.Time `db:"created_at"`
	Updated_at  time.Time `db:"updated_at"`
	Status      uint8     `db:"status"`
}
