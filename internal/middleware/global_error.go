package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/library-api-gin/internal/customerror"
	"github.com/natanaelrusli/library-api-gin/internal/domain"
	"github.com/natanaelrusli/library-api-gin/internal/dto/httpdto"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *customerror.CustomError:
				ctx.JSON(e.Code, httpdto.ErrorResponse{
					Message: e.Message,
					Details: nil,
				})
			default:
				ctx.JSON(500, httpdto.ErrorResponse{
					Message: domain.ErrInternalServerError.Error(),
					Details: nil,
				})
			}
			ctx.Abort()
		}
	}
}
