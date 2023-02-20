package models

import (
	uuid "github.com/satori/go.uuid"
)

type Author struct { // table name: authors
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`

	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	YearBorn    int    `json:"year_born"`
}
