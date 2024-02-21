package postgres

import (
	"context"
	"exam/api/models"
	"exam/config"
	"exam/pkg/logger"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBookRepo_Create(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	pgStore, err := New(context.Background(), cfg, log)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createBook := models.CreateBook{
		Name:       "name 1",
		PageNumber: 2,
	}
	id, err := pgStore.Book().Create(context.Background(), createBook)
	if err != nil {
		t.Error("error while inserting book", err)

	}
	bookID, err := pgStore.Book().GetByID(context.Background(), models.PrimaryKey{
		ID: id,
	})
	if err != nil {
		t.Error("error", err)
	}
	if id == "" {
		t.Error("error while creating book")
	}
	assert.Equal(t, bookID.Name, createBook.Name)
	assert.Equal(t, bookID.PageNumber, createBook.PageNumber)

}

func TestBookRepo_GetByID(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	pgStore, err := New(context.Background(), cfg, log)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}
	books, err := pgStore.Book().GetList(context.Background(), models.GetListRequest{
		Page:   1,
		Limit:  1,
		Search: "",
	})

	if len(books.Books) == 0 {
		t.Error("error", err)

	}

	expectedbooks := books.Books[0].ID
	t.Run("succes", func(t *testing.T) {
		book, err := pgStore.Book().GetByID(context.Background(), models.PrimaryKey{ID: expectedbooks})

		if err != nil {
			t.Error("error while geting by id book", err)
		}
		if book.ID != expectedbooks {
			t.Errorf("expected: %q but got: %q", expectedbooks, book.ID)
		}
		if book.Name == "" {
			t.Error("expected: book name but got : nothing")
		}
		if book.PageNumber <= 0 {

			t.Errorf("expeceted: more than 0 page number but got %q", book.PageNumber)
		}

	})

	t.Run("fail", func(t *testing.T) {
		bookID := ""
		book, err := pgStore.Book().GetByID(context.Background(), models.PrimaryKey{
			ID: bookID,
		})
		if err != nil {
			t.Error("error while getting book id", err)
		}
		if book.ID != bookID {
			t.Errorf("expected: %q, but got %q", bookID, book.ID)
		}
		if book.Name == "" {
			t.Error("expected: book name but got : nothing")
		}
		if book.PageNumber <= 0 {

			t.Errorf("expeceted: more than 0 page number but got %q", book.PageNumber)
		}

	})

}

func TestBookRepo_GetList(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	pgStore, err := New(context.Background(), cfg, log)
	if err != nil {
		t.Errorf("error while connecting db %q", err)
	}

	books, err := pgStore.Book().GetList(context.Background(), models.GetListRequest{
		Page:  1,
		Limit: 1000,
	})
	if err != nil {
		t.Error("error while getting list of book", err.Error())
	}
	if len(books.Books) != 5 {
		t.Errorf("expected 5 rows , but got %q", len(books.Books))
	}

	assert.Equal(t, len(books.Books), 5)

}

func TestBookRepo_Update(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	pgStore, err := New(context.Background(), cfg, log)
	if err != nil {
		t.Error("error while connecting to db ", err)
	}

	createBook := models.CreateBook{
		Name:       "name1",
		PageNumber: 10,
	}

	bookID, err := pgStore.Book().Create(context.Background(), createBook)
	if err != nil {
		t.Error("erro while creating book in tetsing", err)
	}

	if err != nil {
		t.Errorf("error while creating book %v", err)
	}

	updateBook := models.UpdateBook{
		ID:         bookID,
		Name:       "name2",
		PageNumber: 11,
	}

	updatedBookID, err := pgStore.Book().Update(context.Background(), updateBook)
	if err != nil {
		t.Error("error updatinf book in testing", err)
	}

	book, err := pgStore.Book().GetByID(context.Background(), models.PrimaryKey{
		ID: updatedBookID,
	})
	if err != nil {
		t.Error("error while geting by id in testing book", err)
	}
	if updatedBookID == "" {
		t.Error("expected updated book id: but got empty string")
	}

	assert.Equal(t, book.ID, updatedBookID)
	assert.Equal(t, book.Name, updateBook.Name)
	assert.Equal(t, book.PageNumber, updateBook.PageNumber)
}

func TestBookRepo_Delete(t *testing.T) {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	pgStore, err := New(context.Background(), cfg, log)
	if err != nil {
		t.Error("error while connecting to db ", err)
	}

	createBook := models.CreateBook{
		Name:       "name5",
		PageNumber: 10,
	}

	bookID, err := pgStore.Book().Create(context.Background(), createBook)
	if err != nil {
		t.Error("error while creating book in delete testing", err)
	}

	if err = pgStore.Book().Delete(context.Background(), models.PrimaryKey{
		ID: bookID,
	}); err != nil {
		t.Error("error while deleteing book in testing", err)
	}
	t.Run("falied", func(t *testing.T) {
		if bookID == "" {
			t.Error("expected book id but go nothing", err)
		}
	})

}
