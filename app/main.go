package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/natanaelrusli/library-api-gin/internal/config"
	"github.com/natanaelrusli/library-api-gin/internal/dto"
	"github.com/natanaelrusli/library-api-gin/internal/middleware"
	"github.com/natanaelrusli/library-api-gin/internal/pkg/database"

	_authorRepo "github.com/natanaelrusli/library-api-gin/internal/author/repository/postgres"
	_bookRepo "github.com/natanaelrusli/library-api-gin/internal/book/repository/postgres"
	_userRepo "github.com/natanaelrusli/library-api-gin/internal/user/repository/postgres"

	bookUsecase "github.com/natanaelrusli/library-api-gin/internal/book/usecase"
	userUsecase "github.com/natanaelrusli/library-api-gin/internal/user/usecase"

	bookHandler "github.com/natanaelrusli/library-api-gin/internal/book/delivery/http"
	userHandler "github.com/natanaelrusli/library-api-gin/internal/user/delivery/http"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.GlobalErrorHandler())
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}

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

	bookRepository := _bookRepo.NewPostgresBookRepository(db)
	authorRepository := _authorRepo.NewPostgresAuthorRepository(db)
	userRepository := _userRepo.NewPostgresUserRepository(db)

	bookUsecase := bookUsecase.NewBookUsecase(bookRepository, authorRepository)
	userUsecase := userUsecase.NewUserUsecase(userRepository)

	bookHandler := bookHandler.NewBookHandler(bookUsecase)
	userHandler := userHandler.NewUserHandler(userUsecase)

	r.GET("/books", bookHandler.GetAllBooks)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.POST("/books", bookHandler.CreateOne)

	r.GET("/books/:id/author", bookHandler.GetBookAuthor)
	r.GET("/books/author", bookHandler.GetAllBooksWithAuthor)

	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/users/:name", userHandler.GetUserByName)

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
