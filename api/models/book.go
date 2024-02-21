package models

type Book struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AuthorName string `json:"author_name"`
	PageNumber int    `json:"page_number"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  int    `json:"deleted_at"`
}

type CreateBook struct {
	Name       string `json:"name"`
	AuthorName   string `json:"author_name"`
	PageNumber int    `json:"page_number"`
}

type UpdateBook struct {
	ID         string `json:"-"`
	Name       string `json:"name"`
	AuthorName string `json:"author_name"`
	PageNumber int    `json:"page_number"`
}

type UpdateBookName struct {
	ID   string `json:"-"`
	Name string `json:"name"`
}

type UpdateBookPageNumber struct {
	ID         string `json:"-"`
	PageNumber int    `json:"page_number"`
}

type BookResponse struct {
	Books []Book `json:"books"`
	Count int    `json:"count"`
}
