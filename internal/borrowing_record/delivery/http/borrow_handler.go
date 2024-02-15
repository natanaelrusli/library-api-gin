package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/library-api-gin/internal/constants"
	"github.com/natanaelrusli/library-api-gin/internal/customerror"
	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/natanaelrusli/library-api-gin/internal/dto"
)

type BorrowingRecordHandler struct {
	BorrowingRecordUsecase domain.BorrowingRecordUsecase
}

func NewBorrowingRecordHandler(bru domain.BorrowingRecordUsecase) *BorrowingRecordHandler {
	return &BorrowingRecordHandler{
		BorrowingRecordUsecase: bru,
	}
}

func (h *BorrowingRecordHandler) Create(c *gin.Context) {
	var req dto.CreateBorrowingRecordRequest
	err := c.BindJSON(&req)
	if err != nil {
		err := customerror.NewCustomError(500, err.Error())
		c.Error(err)

		return
	}

	record, err := h.BorrowingRecordUsecase.CreateRecord(
		req.UserId,
		req.BookId,
		req.Status,
	)

	if err != nil {
		err := customerror.NewCustomError(http.StatusBadRequest, err.Error())
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": constants.MessageOK,
		"data":    record,
	})
}

func (h *BorrowingRecordHandler) GetAllBorrowed(c *gin.Context) {
	records, err := h.BorrowingRecordUsecase.GetAllBorrowedRecord()
	if err != nil {
		err := customerror.NewCustomError(http.StatusInternalServerError, err.Error())
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Message: constants.MessageOK,
		Data:    records,
	})
}
