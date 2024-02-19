package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/library-api-gin/internal/constants"
	"github.com/natanaelrusli/library-api-gin/internal/customerror"
	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/natanaelrusli/library-api-gin/internal/dto"
	"github.com/natanaelrusli/library-api-gin/internal/dto/httpdto"
)

type BorrowingRecordHandler struct {
	BorrowingRecordUsecase domain.BorrowingRecordUsecase
	BookUsecase            domain.BookUsecase
}

func NewBorrowingRecordHandler(bru domain.BorrowingRecordUsecase, bu domain.BookUsecase) *BorrowingRecordHandler {
	return &BorrowingRecordHandler{
		BorrowingRecordUsecase: bru,
		BookUsecase:            bu,
	}
}

func (h *BorrowingRecordHandler) Create(ctx *gin.Context) {
	var req dto.CreateBorrowingRecordRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		err := customerror.NewCustomError(500, err.Error())
		ctx.Error(err)

		return
	}

	record, err := h.BorrowingRecordUsecase.CreateRecord(
		ctx,
		req.UserId,
		req.BookId,
		req.Status,
	)

	if err != nil {
		err := customerror.NewCustomError(http.StatusBadRequest, err.Error())
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": constants.MessageOK,
		"data":    record,
	})
}

func (h *BorrowingRecordHandler) GetAllBorrowed(ctx *gin.Context) {
	records, err := h.BorrowingRecordUsecase.GetAllBorrowedRecord(ctx)
	if err != nil {
		err := customerror.NewCustomError(http.StatusInternalServerError, err.Error())
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpdto.Response{
		Message: constants.MessageOK,
		Data:    records,
	})
}

func (h *BorrowingRecordHandler) Borrow(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var req dto.UpdateBookStockRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		err := customerror.NewCustomError(http.StatusInternalServerError, err.Error())
		ctx.Error(err)

		return
	}

	book, err := h.BookUsecase.DecreaseStock(ctx, req.BookId, req.Amount)
	if err != nil {
		err := customerror.NewCustomError(http.StatusConflict, err.Error())
		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusOK, httpdto.UpdateStockResponse{
		Id:    book.Id,
		Stock: int(book.Stock),
	})
}
