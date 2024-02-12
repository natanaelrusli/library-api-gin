package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natanaelrusli/library-api-gin/internal/dto"
	"github.com/natanaelrusli/library-api-gin/internal/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())

	r.GET("/ping", func(ctx *gin.Context) {
		var query dto.Query
		if err := ctx.ShouldBindQuery(&query); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, query)
		}

		ctx.JSON(200, dto.Response{
			Message: "PONG!!!",
			Data:    query.Name,
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("error gin")
	}
}
