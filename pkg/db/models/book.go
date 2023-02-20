package models

import (
	uuid "github.com/satori/go.uuid"
)

type Book struct { // table name: books
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`

	// Indexing name and author id, ensuring there's no repeated book with the same name from the same author
	Name     string    `json:"name" gorm:"index:idx_book,unique"`
	Author   *Author   `json:"author"`
	AuthorID uuid.UUID `json:"author_id,omitempty" gorm:"index:idx_book,unique"`

	ReleaseYear int `json:"release_year"`
}
