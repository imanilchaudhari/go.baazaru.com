package models

import (
	"time"
)

// *****************************************************************************
// Gender
// *****************************************************************************

// Gender table contains the information for each Gender
type Gender struct {
	Id         uint32    `db:"id"`
	Name       string    `db:"name"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
	Deleted    uint8     `db:"deleted"`
}
