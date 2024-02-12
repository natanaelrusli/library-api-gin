package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Query struct {
	Name string `form:"name"`
}

func main() {
	r := gin.Default()
	r.Use(Logger())

	r.GET("/ping", func(ctx *gin.Context) {
		var query Query
		if err := ctx.ShouldBindQuery(&query); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, query)
		}

		ctx.JSON(200, Response{
			Message: "PONG!!!",
			Data:    query.Name,
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("error gin")
	}
}
