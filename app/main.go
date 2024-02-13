package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	bookHandler "github.com/natanaelrusli/library-api-gin/internal/book/delivery/http"
	"github.com/natanaelrusli/library-api-gin/internal/book/repository/postgres"
	"github.com/natanaelrusli/library-api-gin/internal/book/usecase"
	"github.com/natanaelrusli/library-api-gin/internal/config"
	"github.com/natanaelrusli/library-api-gin/internal/dto"
	"github.com/natanaelrusli/library-api-gin/internal/middleware"
	"github.com/natanaelrusli/library-api-gin/internal/pkg/database"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.GlobalErrorHandler())
	config := config.InitConfig()
	db, err := database.InitPostgres(config)
	if err != nil {
		log.Fatalln("error connecting to database: ", err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	bookRepository := postgres.NewPostgresBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	bookHandler := bookHandler.NewBookHandler(bookUsecase)

	r.GET("/books", bookHandler.GetAllBooks)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.POST("/books", bookHandler.CreateOne)

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
