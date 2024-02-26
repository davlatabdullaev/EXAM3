package service

import (
	"context"
	"exam/api/models"
	"exam/pkg/logger"
	"exam/storage"
)

type bookService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewBookService(storage storage.IStorage, log logger.ILogger) bookService {
	return bookService{
		storage: storage,
		log:     log,
	}
}

// CREATE

func (b bookService) Create(ctx context.Context, book models.CreateBook) (models.Book, error) {
	b.log.Info("book create service layer", logger.Any("book", book))
	id, err := b.storage.Book().Create(ctx, book)
	if err != nil {
		b.log.Error("error in service layer while creating book", logger.Error(err))
		return models.Book{}, err
	}

	createdBook, err := b.storage.Book().GetByID(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		b.log.Error("error is while getting by id", logger.Error(err))
		return models.Book{}, err
	}

	return createdBook, nil
}

// GET BY ID

func (b bookService) Get(ctx context.Context, id string) (models.Book, error) {
	book, err := b.storage.Book().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		b.log.Error("error in service layer while getting book by id", logger.Error(err))
		return models.Book{}, err
	}

	return book, nil
}

// GET LIST

func (b bookService) GetList(ctx context.Context, request models.GetListRequest) (models.BookResponse, error) {
	b.log.Info("book get list service layer", logger.Any("book", request))

	books, err := b.storage.Book().GetList(ctx, request)
	if err != nil {
		b.log.Error("error in service layer  while getting book list", logger.Error(err))
		return models.BookResponse{}, err
	}

	return books, nil
}

// UPDATE

func (b bookService) Update(ctx context.Context, book models.UpdateBook) (models.Book, error) {
	id, err := b.storage.Book().Update(ctx, book)
	if err != nil {
		b.log.Error("error in service layer while updating book", logger.Error(err))
		return models.Book{}, err
	}

	updatedBook, err := b.storage.Book().GetByID(context.Background(), models.PrimaryKey{ID: id})
	if err != nil {
		b.log.Error("error in service layer while getting book by id", logger.Error(err))
		return models.Book{}, err
	}

	return updatedBook, nil
}

// DELETE

func (b bookService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := b.storage.Book().Delete(ctx, key)

	return err
}

// UPDATE BOOK PAGE NUMBER

func (b bookService) UpdateBookPageNumber(ctx context.Context, request models.UpdateBookPageNumber) error {

	err := b.storage.Book().UpdateBookPageNumber(ctx, request)

	if err != nil {
		b.log.Error("error in service layer  while updating book page number", logger.Error(err))
		return err
	}

	return nil
}

// UPDATE BOOK NAME

func (b bookService) UpdateBookName(ctx context.Context, request models.UpdateBookName) error {

	err := b.storage.Book().UpdateBookName(ctx, request)

	if err != nil {
		b.log.Error("error in service layer  while updating book name", logger.Error(err))
		return err
	}

	return nil

}
