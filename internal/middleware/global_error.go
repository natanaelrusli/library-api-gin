package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/library-api-gin/internal/customerror"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *customerror.CustomError:
				ctx.JSON(e.Code, gin.H{
					"error": e.Message,
				})
			default:
				ctx.JSON(500, gin.H{
					"error": err.Error(),
				})
			}
			ctx.Abort()
		}
	}
}
