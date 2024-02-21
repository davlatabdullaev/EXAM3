package postgres

import (
	"context"
	"database/sql"
	"exam/api/models"
	"exam/pkg/logger"
	"exam/storage"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewBookRepo(db *pgxpool.Pool, log logger.ILogger) storage.IBookStorage {
	return &bookRepo{
		db:  db,
		log: log,
	}
}

// CREATE

func (b *bookRepo) Create(ctx context.Context, book models.CreateBook) (string, error) {

	id := uuid.New()

	query := `insert into books 
	(id, 
	name, 
	author_name, 
	page_number) values ($1, $2, $3, $4)`
	if rowsAffected, err := b.db.Exec(ctx, query, 
		id, 
		book.Name, 
		book.AuthorName, 
		book.PageNumber); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error is in rows affected", logger.Error(err))
			return "", err
		}
		b.log.Error("error is while inserting book data", logger.Error(err))
		return "", err
	}

	return id.String(), nil
}

// GET BY ID

func (b *bookRepo) GetByID(ctx context.Context, key models.PrimaryKey) (models.Book, error) {

	var createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	book := models.Book{}

	query := `select 
	id,
	name, 
	author_name,
	page_number,
    created_at,
	updated_at 
	from books where id = $1 and deleted_at = 0 `

	if err := b.db.QueryRow(ctx, query,
		key.ID).Scan(
		&book.ID,
		&book.Name,
		&book.AuthorName,
		&book.PageNumber,
		&createdAt,
		&updatedAt,
	); err != nil {
		b.log.Error("error is while selecting book", logger.Error(err))
		return models.Book{}, err
	}

	if createdAt.Valid {
		book.CreatedAt = createdAt.String
	}

	if updatedAt.Valid {
		book.UpdatedAt = updatedAt.String
	}

	return book, nil

}

// GET LIST

func (b *bookRepo) GetList(ctx context.Context, req models.GetListRequest) (models.BookResponse, error) {

	var (
		books                = []models.Book{}
		count                = 0
		query, countQuery    string
		page                 = req.Page
		offset               = (page - 1) * req.Limit
		search               = req.Search
		createdAt, updatedAt = sql.NullString{}, sql.NullString{}
	)

	countQuery = `select count(1) from books where deleted_at = 0 `

	if search != "" {
		countQuery += fmt.Sprintf(` and name ilike '%%%s%%'`, search)
	}
	if err := b.db.QueryRow(ctx, countQuery).Scan(&count); err != nil {
		b.log.Error("error is while selecting count", logger.Error(err))
		return models.BookResponse{}, err
	}

	query = `select 
	id, 
	name, 
	author_name,
	page_number,
	created_at,
    updated_at 
	from books where deleted_at = 0 `

	if search != "" {
		query += fmt.Sprintf(` and name ilike '%%%s%%'`, search)
	}

	query += ` order by created_at desc LIMIT $1 OFFSET $2`
	rows, err := b.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		b.log.Error("error is while selecting books", logger.Error(err))
		return models.BookResponse{}, err
	}

	for rows.Next() {
		book := models.Book{}
		if err = rows.Scan(
			&book.ID,
			&book.Name,
			&book.AuthorName,
			&book.PageNumber,
			&createdAt,
			&updatedAt); err != nil {
			b.log.Error("error is while scanning books data", logger.Error(err))
			return models.BookResponse{}, err
		}

		if createdAt.Valid {
			book.CreatedAt = createdAt.String
		}

		if updatedAt.Valid {
			book.UpdatedAt = updatedAt.String
		}

		books = append(books, book)

	}

	return models.BookResponse{
		Books: books,
		Count: count,
	}, nil
}

// UPDATE

func (b *bookRepo) Update(ctx context.Context, request models.UpdateBook) (string, error) {

	book := models.Book{}

	query := `update books set 
	name = $1, 
	author_name = $2, 
	page_number = $3,
	updated_at = now() 
	where id = $4`

	if rowsAffected, err := b.db.Exec(ctx, query,

		&request.Name,
		&request.AuthorName,
		&request.PageNumber,
		&request.ID,
	); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error is in rows affected", logger.Error(err))
			return "", err
		}
		return book.ID, err
	}

	getQuery := `select 
	id, 
	name, 
	author_name,
	page_number
	from books where id = $1`

	if err := b.db.QueryRow(ctx, getQuery, request.ID).Scan(
		&book.ID,
		&book.Name,
		&book.AuthorName,
		&book.PageNumber); err != nil {
		b.log.Error("error is while selecting book after update", logger.Error(err))
		return "", err
	}
	return book.ID, nil

}

// DELETE

func (b *bookRepo) Delete(ctx context.Context, key models.PrimaryKey) error {

	query := `update books set 
	deleted_at = extract(epoch from current_timestamp) 
	   where id = $1`
	if rowsAffected, err := b.db.Exec(ctx, query, key.ID); err != nil {
		if r := rowsAffected.RowsAffected(); r == 0 {
			b.log.Error("error is while deleting book", logger.Error(err))
			return err
		}
		return err
	}
	return nil
}

// UPDATE BOOK PAGE NUMBER

func (b *bookRepo) UpdateBookPageNumber(ctx context.Context, request models.UpdateBookPageNumber) error {

	query := `
	update books
			set 
			page_number = $1, 
			updated_at = now()
				where id = $2 `

	if _, err := b.db.Exec(ctx, query, request.PageNumber, request.ID); err != nil {
		fmt.Println("error while updating page number for book", err.Error())
		return err
	}

	return nil

}

// UPDATE BOOK NAME

func (b *bookRepo) UpdateBookName(ctx context.Context, request models.UpdateBookName) error {

	query := `
	update books
			set 
			name = $1, 
			updated_at = now()
				where id = $2 `

	if _, err := b.db.Exec(ctx, query, request.Name, request.ID); err != nil {
		fmt.Println("error while updating name for book", err.Error())
		return err
	}

	return nil

}
