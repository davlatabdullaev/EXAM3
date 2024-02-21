package storage

import (
	"context"
	"exam/api/models"
)

type IStorage interface {
	Close()
	Book() IBookStorage
	//Author() IAuthorStorage
}

type IBookStorage interface {
	Create(context.Context, models.CreateBook) (string, error)
	GetByID(context.Context, models.PrimaryKey) (models.Book, error)
	GetList(context.Context, models.GetListRequest) (models.BookResponse, error)
	Update(context.Context, models.UpdateBook) (string, error)
	Delete(context.Context, models.PrimaryKey) error
	UpdateBookName(context.Context, models.UpdateBookName) error
	UpdateBookPageNumber(context.Context, models.UpdateBookPageNumber) error
}

// type IAuthorStorage interface {
// 	Create(context.Context, models.CreateAuthor) (string, error)
// 	Get(context.Context, models.PrimaryKey) (models.Author, error)
// 	GetList(context.Context, models.GetListRequest) (models.AuthorsResponse, error)
// 	Update(context.Context, models.UpdateAuthor) (string, error)
// 	Delete(context.Context, string) error
// 	UpdatePassword(context.Context, models.UpdateAuthorPassword) error
// 	GetPassword(context.Context, string) (string, error)
// }
