package httpserver

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/natanaelrusli/library-api-gin/internal/config"
	"github.com/natanaelrusli/library-api-gin/internal/dto"
	"github.com/natanaelrusli/library-api-gin/internal/dto/httpdto"
	"github.com/natanaelrusli/library-api-gin/internal/middleware"
	"github.com/natanaelrusli/library-api-gin/internal/pkg/database"

	_authorRepo "github.com/natanaelrusli/library-api-gin/internal/author/repository/postgres"
	_bookRepo "github.com/natanaelrusli/library-api-gin/internal/book/repository/postgres"
	_borrowingRecordRepo "github.com/natanaelrusli/library-api-gin/internal/borrowing_record/repository/postgres"
	_userRepo "github.com/natanaelrusli/library-api-gin/internal/user/repository/postgres"

	bookUsecase "github.com/natanaelrusli/library-api-gin/internal/book/usecase"
	borrowingRecordUsecase "github.com/natanaelrusli/library-api-gin/internal/borrowing_record/usecase"
	userUsecase "github.com/natanaelrusli/library-api-gin/internal/user/usecase"

	bookHandler "github.com/natanaelrusli/library-api-gin/internal/book/delivery/http"
	borrowingRecordHandler "github.com/natanaelrusli/library-api-gin/internal/borrowing_record/delivery/http"
	userHandler "github.com/natanaelrusli/library-api-gin/internal/user/delivery/http"
)

func initServer(config *config.Config) *http.Server {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.GlobalErrorHandler())
	r.ContextWithFallback = true

	db, err := database.InitPostgres(config)

	if err != nil {
		log.Fatalln("error connecting to database: ", err)
	}

	bookRepository := _bookRepo.NewPostgresBookRepository(db)
	authorRepository := _authorRepo.NewPostgresAuthorRepository(db)
	userRepository := _userRepo.NewPostgresUserRepository(db)
	borrowingRecordRepository := _borrowingRecordRepo.NewBorrowingRecordRepository(db)

	bookUsecase := bookUsecase.NewBookUsecase(bookRepository, authorRepository, borrowingRecordRepository)
	userUsecase := userUsecase.NewUserUsecase(userRepository)
	borrowingRecordUsecase := borrowingRecordUsecase.NewBorrowingRecordUsecase(borrowingRecordRepository, bookRepository)

	bookHandler := bookHandler.NewBookHandler(bookUsecase)
	userHandler := userHandler.NewUserHandler(userUsecase)
	borrowingRecordHandler := borrowingRecordHandler.NewBorrowingRecordHandler(borrowingRecordUsecase, bookUsecase)

	r.GET("/books", bookHandler.GetAllBooks)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.POST("/books", bookHandler.CreateOne)

	r.GET("/books/:id/author", bookHandler.GetBookAuthor)
	r.GET("/books/author", bookHandler.GetAllBooksWithAuthor)

	r.GET("/users", userHandler.GetUsers)

	r.GET("/borrowing-records/borrowed", borrowingRecordHandler.GetAllBorrowed)
	r.POST("/borrowing-records", borrowingRecordHandler.Create)

	r.POST("/borrow", borrowingRecordHandler.Borrow)

	r.GET("/ping", func(ctx *gin.Context) {
		var query dto.Query
		if err := ctx.ShouldBindQuery(&query); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, query)
		}

		ctx.JSON(200, httpdto.Response{
			Message: "PONG!!!",
			Data:    query.Name,
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("error gin")
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return srv
}

func StartGinServer() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}

	srv := initServer(config)

	// graceful shutdown
	go func() {
		log.Println("running server on port :", ":8080")
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatal("error while server listen and serve: ", err)
			}
		}
		log.Println("server is not receiving new requests...")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	graceDuration := time.Duration(5) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), graceDuration)
	defer cancel()

	log.Println("attempt to shutting down the server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("error shutting down server: ", err)
	}

	log.Println("http server is shutting down gracefully")
}
