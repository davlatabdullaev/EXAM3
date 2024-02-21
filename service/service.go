package service

import (
	"exam/pkg/logger"
	"exam/storage"
)

type IServiceManager interface {
	Book() bookService
//	Author() authorService
}

type Service struct {
	bookService bookService
//	authorService authorService
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	services := Service{}

	services.bookService = NewBookService(storage, log)

	return services
}


func (s Service) Book() bookService {
	return s.bookService
}

// func (s Service) Author() authorService {
// 	return s.authorService
// }