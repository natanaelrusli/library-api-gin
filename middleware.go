package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Set("example", "12345")
		ctx.Next()

		latency := time.Since(t)
		log.Println(latency)

		status := ctx.Writer.Status()
		log.Println(status)
	}
}
