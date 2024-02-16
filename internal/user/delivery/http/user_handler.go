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

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	name, exists := ctx.GetQuery("name")
	if !exists {
		users, err := h.UserUsecase.FetchAll(ctx)
		if err != nil {
			err := customerror.NewCustomError(http.StatusInternalServerError, err.Error())
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dto.Response{
			Message: constants.MessageOK,
			Data:    users,
		})
		return
	} else {
		if !exists {
			err := customerror.NewCustomError(http.StatusBadRequest, domain.ErrBadParamInput.Error())
			ctx.Error(err)
			return
		}

		user, err := h.UserUsecase.FetchByName(ctx, name)
		if err != nil {
			err := customerror.NewCustomError(http.StatusInternalServerError, err.Error())
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dto.Response{
			Message: constants.MessageOK,
			Data:    user,
		})
	}

}
