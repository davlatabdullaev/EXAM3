 package service

// import (
// 	"context"
// 	"errors"
// 	"exam/api/models"
// 	"exam/pkg/check"
// 	"exam/pkg/logger"
// 	"exam/storage"

// 	"github.com/jackc/pgx/v5"
// )

// type authorService struct {
// 	storage storage.IStorage
// 	log     logger.ILogger
// }

// func NewAuthorService(storage storage.IStorage, log logger.ILogger) authorService {
// 	return authorService{
// 		storage: storage,
// 		log:     log,
// 	}
// }

// // CREATE

// func (a authorService) Create(ctx context.Context, createAuthor models.CreateAuthor) (models.Author, error) {
//     a.log.Info("author create service layer", logger.Any("author", createAuthor))
// 	pKey, err := a.storage.Author().Create(ctx, createAuthor)
// 	if err != nil {
// 		a.log.Error("error in service layer while creating author ", logger.Error(err))
// 		return models.Author{}, err
// 	}

// 	author, err := a.storage.Author().Get(ctx, models.PrimaryKey{
// 		ID: pKey,
// 	})
// 	if err != nil {
// 		a.log.Error("error in service layer while getting author", logger.Error(err))
// 		return models.Author{}, err
// 	}

// 	return author, nil
// }

// // GET

// func (a authorService) Get(ctx context.Context, pkey models.PrimaryKey) (models.Author, error) {

// 	author, err := a.storage.Author().Get(ctx, pkey)
// 	if err != nil {
// 		if !errors.Is(err, pgx.ErrNoRows) {
// 			a.log.Error("error in service layer while getting author by id", logger.Error(err))
// 			return models.Author{}, err
// 		}
// 	}

// 	return author, nil
// }

// // GET LIST

// func (a authorService) GetList(ctx context.Context, request models.GetListRequest) (models.AuthorsResponse, error) {

// 	authors, err := a.storage.Author().GetList(ctx, request)
// 	if err != nil {
// 		if !errors.Is(err, pgx.ErrNoRows) {
// 			a.log.Error("error in service layer while getting authors list", logger.Error(err))
// 			return models.AuthorsResponse{}, err
// 		}
// 	}
// 	return authors, nil
// }

// // UPDATE

// func (a authorService) Update(ctx context.Context, updateAuthor models.UpdateAuthor) (models.Author, error) {

// 	id, err := a.storage.Author().Update(ctx, updateAuthor)
// 	if err != nil {
// 		a.log.Error("error in servise layer updating author by id", logger.Error(err))
// 		return models.Author{}, err
// 	}

// 	author, err := a.storage.Author().Get(ctx, models.PrimaryKey{
// 		ID: id,
// 	})
// 	if err != nil {
// 		a.log.Error("error in service layer getting author after update", logger.Error(err))
// 		return models.Author{}, err
// 	}

// 	return author, nil
// }

// // DELETE

// func (a authorService) Delete(ctx context.Context, id string) error {

// 	err := a.storage.Author().Delete(ctx, id)

// 	return err
// }

// // UPDATE PASSWORD

// func (a authorService) UpdatePassword(ctx context.Context, request models.UpdateAuthorPassword) error {

// 	oldPassword, err := a.storage.Author().GetPassword(ctx, request.ID)
// 	if err != nil {
// 		a.log.Error("error in service layer getting password by id", logger.Error(err))
// 		return err
// 	}

// 	if oldPassword != request.OldPassword {
// 		a.log.Error("error in service layer old password is not correct")
// 		return errors.New("old password did not match")
// 	}

// 	if err = check.ValidatePassword(request.NewPassword); err != nil {
// 		a.log.Error("error in service layer new password validation failed", logger.Error(err))
// 		return err
// 	}

// 	if err = a.storage.Author().UpdatePassword(context.Background(), request); err != nil {
// 		a.log.Error("error in service layer while updating author password ", logger.Error(err))
// 		return err
// 	}

// 	return nil
// }
