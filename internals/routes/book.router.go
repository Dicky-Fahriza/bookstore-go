package routes

import (
	"bookstore-go/internals/handlers"
	"bookstore-go/internals/middlewares"
	"bookstore-go/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitBookRouter(router *gin.Engine, db *sqlx.DB) {
	bookRouter := router.Group("/book")
	bookRepo := repositories.InitBookRepo(db)
	BookHandler := handlers.InitBookHandler(bookRepo)

	// localhost:8000/book
	bookRouter.GET("", middlewares.CheckToken, BookHandler.GetBooks)

	// get book by id
	bookRouter.GET("/:id", middlewares.CheckToken, BookHandler.GetBookById)

	// localhost:8000/book/new
	bookRouter.POST("/new", middlewares.CheckToken, BookHandler.CreateBooks)

	// update book by id
	bookRouter.PATCH("/:id", middlewares.CheckToken, BookHandler.UpdateBookById)

	// localhost:8000/book/:id
	router.DELETE("/books/:id", middlewares.CheckToken, BookHandler.DeleteBook)

}
