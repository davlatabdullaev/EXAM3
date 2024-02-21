package handler

import (
	"context"
	"exam/api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CREATE

// CreateBook godoc
// @Router       /book [POST]
// @Summary      Creates a new book
// @Description  create a new book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        book body models.CreateBook false "book"
// @Success      201  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBook(c *gin.Context) {
	createBook := models.CreateBook{}

	if err := c.ShouldBindJSON(&createBook); err != nil {
		handleResponse(c, h.log, "error is while decoding for book", http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Book().Create(context.Background(), createBook)
	if err != nil {
		handleResponse(c, h.log, "error is while creating book", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.log, "", http.StatusCreated, res)
}

// GET BY ID

// GetBook godoc
// @Router       /book/{id} [GET]
// @Summary      Get book by id
// @Description  get book by id
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id path string true "book_id"
// @Success      201  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBook(c *gin.Context) {
	var err error

	uid := c.Param("id")

	book, err := h.services.Book().Get(context.Background(), uid)
	if err != nil {
		handleResponse(c, h.log, "error is while getting book by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, book)
}

// GET LIST

// GetBookList godoc
// @Router       /books [GET]
// @Summary      Get book list
// @Description  get book list
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201  {object}  models.BookResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBookList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting pageStr", http.StatusBadRequest, err)
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, h.log, "error is while converting limitStr", http.StatusBadRequest, err)
		return
	}

	search = c.Query("search")

	baskets, err := h.services.Book().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, h.log, "error is while getting book list", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, baskets)
}

//UPDATE

// UpdateBook godoc
// @Router       /book/{id} [PUT]
// @Summary      Update book
// @Description  update book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id path string true "book_id"
// @Param        book body models.UpdateBook false "book"
// @Success      201  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBook(c *gin.Context) {
	updatedBook := models.UpdateBook{}

	uid := c.Param("id")
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		handleResponse(c, h.log, "error is while decoding ", http.StatusBadRequest, err)
		return
	}

	updatedBook.ID = uid

	book, err := h.services.Book().Update(context.Background(), updatedBook)
	if err != nil {
		handleResponse(c, h.log, "error is while updating book", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, book)
}

// DELETE

// DeleteBook godoc
// @Router       /book/{id} [Delete]
// @Summary      Delete book
// @Description  delete book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id path string true "book_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBook(c *gin.Context) {
	uid := c.Param("id")

	if err := h.services.Book().Delete(context.Background(), models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, h.log, "error is while deleting book", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, nil)
}

// UPDATE BOOK PAGE NUMBER

// UpdateBookPageNumber godoc
// @Router       /book/{id} [PATCH]
// @Summary      Update book page number
// @Description  update book page number
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Param        Book body models.UpdateBookPageNumber true "book"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBookPageNumber(c *gin.Context) {
	updateBookPageNumber := models.UpdateBookPageNumber{}

	if err := c.ShouldBindJSON(&updateBookPageNumber); err != nil {
		handleResponse(c, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, h.log, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateBookPageNumber.ID = uid.String()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = h.services.Book().UpdateBookPageNumber(ctx, updateBookPageNumber); err != nil {
		handleResponse(c, h.log, "error while updating book page number", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "page number successfully updated")
}


// UPDATE BOOK NAME


// UpdateBookName godoc
// @Router       /book_name/{id} [PATCH]
// @Summary      Update book name
// @Description  update book name
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Param        Book body models.UpdateBookName true "book"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBookName(c *gin.Context) {
	updateBookName := models.UpdateBookName{}

	if err := c.ShouldBindJSON(&updateBookName); err != nil {
		handleResponse(c, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, h.log, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateBookName.ID = uid.String()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = h.services.Book().UpdateBookName(ctx, updateBookName); err != nil {
		handleResponse(c, h.log, "error while updating book name", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "book name successfully updated")
}
