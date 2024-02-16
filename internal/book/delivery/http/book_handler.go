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

func (h *BookHandler) GetAllBooks(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	books, err := h.BookUsecase.FetchAll(ctx)
	if err != nil {
		err := customerror.NewCustomError(400, err.Error())
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: constants.MessageOK,
		Data:    books,
	})
}

func (h *BookHandler) GetBookByID(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var params dto.GetBookByIdParams

	if err := ctx.ShouldBindUri(&params); err != nil {
		err := customerror.NewCustomError(400, err.Error())
		ctx.Error(err)
		return
	}

	book, err := h.BookUsecase.GetByID(ctx, params.ID)
	if err != nil {
		err := customerror.NewCustomError(500, err.Error())
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    book,
	})
}

func (h *BookHandler) CreateOne(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var req dto.CreateOneBookRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		err := customerror.NewCustomError(500, err.Error())
		ctx.Error(err)

		return
	}

	book, err := h.BookUsecase.CreateOne(
		ctx,
		req.Title,
		req.Description,
		req.Cover,
		req.AuthorId,
		req.Stock,
	)

	if err != nil {
		err := customerror.NewCustomError(500, err.Error())
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    book,
	})
}

func (h *BookHandler) GetBookAuthor(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var params dto.GetAuthorByBookIdParams

	if err := ctx.ShouldBindUri(&params); err != nil {
		err := customerror.NewCustomError(400, err.Error())
		ctx.Error(err)
		return
	}

	author, err := h.BookUsecase.GetBookAuthor(ctx, params.ID)
	if err != nil {
		err := customerror.NewCustomError(500, err.Error())
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    author,
	})
}

func (h *BookHandler) GetAllBooksWithAuthor(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	books, err := h.BookUsecase.FetchAllWithAuthor(ctx)
	if err != nil {
		err := customerror.NewCustomError(400, err.Error())
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Message: constants.MessageOK,
		Data:    books,
	})
}
