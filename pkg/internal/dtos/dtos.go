package dtos

type CreateBookBody struct {
	Name        string `json:"name" validate:"required"`
	AuthorID    string `json:"author_id" validate:"required"`
	ReleaseYear int    `json:"released_year" validate:"required"`
}

type CreateAuthorBody struct {
	Name        string `json:"name" validate:"required"`
	Nationality string `json:"nationality" validate:"required"`
	YearBorn    int    `json:"year_born" validate:"required"`
}
