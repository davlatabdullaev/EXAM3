package handler

import (
	"context"
	"errors"
	"net/http"
	"exam/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CREATE

// CreateAuthor godoc
// @Router       /author [POST]
// @Summary      Create a new author
// @Description  Create a new author
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        author  body  models.CreateAuthor  true  "author data"
// @Success      201  {object}  models.Author
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateAuthor(c *gin.Context) {
	createAuthor := models.CreateAuthor{}

	if err := c.ShouldBindJSON(&createAuthor); err != nil {
		handleResponse(c,  h.log, "error while reading body from client", http.StatusBadRequest, err)
		return
	}

	resp, err := h.services.Author().Create(context.Background(), createAuthor)
	if err != nil {
		handleResponse(c,  h.log, "error while creating author", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c,  h.log, "", http.StatusCreated, resp)

}

// GET  BY ID

// GetAuthorByID godoc
// @Router       /author/{id} [GET]
// @Summary      Get author by id
// @Description  Get auhtor by id
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        id path string true "author"
// @Success      200  {object}  models.Author
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetAuthorByID(c *gin.Context) {

	var err error

	uid := c.Param("id")

	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c,  h.log, "invalid uuid type ", http.StatusBadRequest, err.Error())
		return
	}

	author, err := h.services.Author().Get(context.Background(), models.PrimaryKey{
		ID: id.String(),
	})
	if err != nil {
		handleResponse(c,  h.log, "error while get author by id", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c,  h.log, "", http.StatusOK, author)

}

// GET LIST

// GetAuthorsList godoc
// @Router       /authors [GET]
// @Summary      Get authors list
// @Description  Get authors list
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      200  {object}  models.AuthorsResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetAuthorList(c *gin.Context) {

	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c,  h.log, "error while parsing page ", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c,  h.log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")
	response, err := h.services.Author().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c,  h.log, "error while getting author", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c,  h.log, "", http.StatusOK, response)

}

// UPDATE 

// UpdateAuthor godoc
// @Router       /author/{id} [PUT]
// @Summary      Update author by id
// @Description  Update author by id
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        id path string true "author id"
// @Param        author body models.UpdateAuthor true "author"
// @Success      200  {object}  models.Author
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateAuthor(c *gin.Context) {
	updateAuthor := models.UpdateAuthor{}

	uid := c.Param("id")
	if uid == "" {
		handleResponse(c,  h.log, "invalid uuid", http.StatusBadRequest, errors.New("uuid is not valid"))
		return
	}

	updateAuthor.ID = uid

	if err := c.ShouldBindJSON(&updateAuthor); err != nil {
		handleResponse(c,  h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	author, err := h.services.Author().Update(context.Background(), updateAuthor)
	if err != nil {
		handleResponse(c,  h.log, "error while updating author", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c,  h.log, "", http.StatusOK, author)

}

// DELETE

// DeleteAuthor godoc
// @Router       /author/{id} [DELETE]
// @Summary      Delete Author
// @Description  Delete Author
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        id path string true "author id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteAuthor(c *gin.Context) {

	uid := c.Param("id")
	id, err := uuid.Parse(uid)
	if err != nil {
		handleResponse(c,  h.log, "uuid is not valid", http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Author().Delete(context.Background(), id.String()); err != nil {
		handleResponse(c,  h.log, "error while deleting author by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c,  h.log, "", http.StatusOK, "data succesfully deleted")

}

// UPDATE PASSWORD

// UpdateAuthorPassword godoc
// @Router       /author/{id} [PATCH]
// @Summary      Update author password
// @Description  update author password
// @Tags         author
// @Accept       json
// @Produce      json
// @Param 		 id path string true "author_id"
// @Param        author body models.UpdateAuthorPassword true "author"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateAuthorPassword(c *gin.Context) {
	updateAuthorPassword := models.UpdateAuthorPassword{}

	if err := c.ShouldBindJSON(&updateAuthorPassword); err != nil {
		handleResponse(c,  h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c,  h.log, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateAuthorPassword.ID = uid.String()

	err = h.services.Author().UpdatePassword(context.Background(), updateAuthorPassword)
	if err != nil {
		handleResponse(c,  h.log, "error while updating author by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.log, "", http.StatusOK, "password successfully updated")
}
