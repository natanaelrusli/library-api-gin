package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/library-api-gin/internal/domain"
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
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
