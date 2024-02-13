package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/library-api-gin/internal/book/dto"
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
		"message": "OK",
		"data":    books,
	})
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var params dto.GetBookByIdParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	book, err := h.BookUsecase.GetByID(params.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    book,
	})
}
