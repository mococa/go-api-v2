package modules

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mococa/go-api-v2/pkg/db/models"
	"github.com/mococa/go-api-v2/pkg/internal/dtos"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

/* -------------- Types -------------- */

type BookModule struct {
	Db *gorm.DB
}

/* -------------- Initilization -------------- */

func NewBookModule(db *gorm.DB) BookModule {
	return BookModule{db}
}

/* -------------- Methods -------------- */

func (module *BookModule) CreateBook(book *dtos.CreateBookBody) (*models.Book, error) {
	validate := validator.New()
	err := validate.Struct(book)
	if err != nil {
		return nil, errors.New("invalid body")
	}

	created_book := models.Book{}

	// Open the directory that contains the file
	fileSystem := os.DirFS("./pkg/db/queries")

	// Read the contents of the file
	query, err := fs.ReadFile(fileSystem, "book.create.sql")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	err = module.Db.Raw(string(query),
		book.Name,
		book.AuthorID,
		book.ReleaseYear,
	).Scan(&created_book).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("book already created")
		}

		return nil, fmt.Errorf("failed to scan book: %s", err.Error())
	}

	if created_book.ID.String() == "" || created_book.AuthorID == uuid.Nil {
		return nil, errors.New("failed create book")
	}

	return &created_book, nil
}

func (module *BookModule) ListBooks() (*[]models.Book, error) {
	books := []models.Book{}

	// Open the directory that contains the file
	fileSystem := os.DirFS("./pkg/db/queries")

	// Read the contents of the file
	query, err := fs.ReadFile(fileSystem, "book.list.sql")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	rows, err := module.Db.Raw(string(query)).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		b := models.Book{}

		err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.AuthorID,
			&b.ReleaseYear,
		)

		if err != nil {
			panic("Failed to scan rows!")
		}

		books = append(books, b)
	}

	return &books, nil
}

func (module *BookModule) FindBook(book_id string) (*models.Book, error) {
	missing_book := book_id == "" || uuid.FromStringOrNil(book_id) == uuid.Nil

	if missing_book {
		return nil, errors.New("invalid book id")
	}

	book := models.Book{}

	// Open the directory that contains the file
	fileSystem := os.DirFS("./pkg/db/queries")

	// Read the contents of the file
	query, err := fs.ReadFile(fileSystem, "book.find.sql")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	err = module.Db.Raw(string(query), book_id).Scan(&book).Error

	if err != nil {
		return nil, fmt.Errorf("failed to scan book: %s", err.Error())
	}

	return &book, nil
}
