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

type AuthorModule struct {
	Db *gorm.DB
}

/* -------------- Initilization -------------- */

func NewAuthorModule(db *gorm.DB) AuthorModule {
	return AuthorModule{db}
}

/* -------------- Methods -------------- */

func (module *AuthorModule) CreateAuthor(author *dtos.CreateAuthorBody) (*models.Author, error) {
	validate := validator.New()
	err := validate.Struct(author)
	if err != nil {
		return nil, errors.New("invalid body")
	}

	created_author := models.Author{}

	// Open the directory that contains the file
	fileSystem := os.DirFS("./pkg/db/queries")

	// Read the contents of the file
	query, err := fs.ReadFile(fileSystem, "author.create.sql")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	err = module.Db.Raw(string(query),
		author.Name,
		author.Nationality,
		author.YearBorn,
	).Scan(&created_author).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("author already created")
		}

		return nil, fmt.Errorf("failed to scan book: %s", err.Error())
	}

	if created_author.ID.String() == "" || created_author.ID == uuid.Nil {
		return nil, errors.New("failed create book")
	}

	return &created_author, nil
}

func (module *AuthorModule) ListAuthorBooks(author_id string) (*[]models.Book, error) {
	missing_author := author_id == "" || uuid.FromStringOrNil(author_id) == uuid.Nil

	if missing_author {
		return nil, errors.New("invalid author id")
	}

	books := []models.Book{}

	// Open the directory that contains the file
	fileSystem := os.DirFS("./pkg/db/queries")

	// Read the contents of the file
	query, err := fs.ReadFile(fileSystem, "book.list_by_author.sql")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	rows, err := module.Db.Raw(string(query), author_id).Rows()
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

func (module *AuthorModule) ListAuthors() (*[]models.Author, error) {
	authors := []models.Author{}

	// Open the directory that contains the file
	fileSystem := os.DirFS("./pkg/db/queries")

	// Read the contents of the file
	query, err := fs.ReadFile(fileSystem, "author.list.sql")
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
		a := models.Author{}

		err := rows.Scan(
			&a.ID,
			&a.Name,
			&a.Nationality,
			&a.YearBorn,
		)

		if err != nil {
			panic("Failed to scan rows!")
		}

		authors = append(authors, a)
	}

	return &authors, nil
}

func (module *AuthorModule) FindAuthor(author_id string) (*models.Author, error) {
	missing_author := author_id == "" || uuid.FromStringOrNil(author_id) == uuid.Nil

	if missing_author {
		return nil, errors.New("invalid author id")
	}

	author := models.Author{}

	// Open the directory that contains the file
	fileSystem := os.DirFS("./pkg/db/queries")

	// Read the contents of the file
	query, err := fs.ReadFile(fileSystem, "author.find.sql")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	err = module.Db.Raw(string(query), author_id).Scan(&author).Error

	if err != nil {
		return nil, fmt.Errorf("failed to scan author: %s", err.Error())
	}

	return &author, nil
}
