package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/library-api-gin/internal/constants"
	"github.com/natanaelrusli/library-api-gin/internal/customerror"
	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/natanaelrusli/library-api-gin/internal/dto"
)

type BookHandler struct {
	BookUsecase domain.BookUsecase
}

func NewBookHandler(u domain.BookUsecase) *BookHandler {
	return &BookHandler{
		BookUsecase: u,
	}
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	books, err := h.BookUsecase.FetchAll()
	if err != nil {
		err := customerror.NewCustomError(400, err.Error())
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Message: constants.MessageOK,
		Data:    books,
	})
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var params dto.GetBookByIdParams

	if err := c.ShouldBindUri(&params); err != nil {
		err := customerror.NewCustomError(400, err.Error())
		c.Error(err)
		return
	}

	book, err := h.BookUsecase.GetByID(params.ID)
	if err != nil {
		err := customerror.NewCustomError(500, err.Error())
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    book,
	})
}

func (h *BookHandler) CreateOne(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var req dto.CreateOneBookRequest

	err := c.BindJSON(&req)
	if err != nil {
		err := customerror.NewCustomError(500, err.Error())
		c.Error(err)

		return
	}

	book, err := h.BookUsecase.CreateOne(
		req.Title,
		req.Description,
		req.Cover,
		req.AuthorId,
		req.Stock,
	)

	if err != nil {
		err := customerror.NewCustomError(500, err.Error())
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    book,
	})
}

func (h *BookHandler) GetBookAuthor(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var params dto.GetAuthorByBookIdParams

	if err := c.ShouldBindUri(&params); err != nil {
		err := customerror.NewCustomError(400, err.Error())
		c.Error(err)
		return
	}

	author, err := h.BookUsecase.GetBookAuthor(params.ID)
	if err != nil {
		err := customerror.NewCustomError(500, err.Error())
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    author,
	})
}
