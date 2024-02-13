package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/library-api-gin/internal/constants"
	"github.com/natanaelrusli/library-api-gin/internal/customerror"
	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/natanaelrusli/library-api-gin/internal/dto"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(u domain.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: u,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	users, err := h.UserUsecase.FetchAll()
	if err != nil {
		err := customerror.NewCustomError(http.StatusInternalServerError, err.Error())
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Message: constants.MessageOK,
		Data:    users,
	})
}
