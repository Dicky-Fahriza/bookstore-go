package routes

import (
	"bookstore-go/internals/handlers"
	"bookstore-go/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAuthRouter(router *gin.Engine, db *sqlx.DB) {
	// Bikin Subrouter
	authRouter := router.Group("/auth")
	authRepo := repositories.InitAuthRepo(db)

	authHandler := handlers.InitAuthHandler(authRepo)

	// Bikin Rute
	// localhost:8000/auth/new
	authRouter.POST("/new", authHandler.Register)
	// localhost:8000/auth/
	authRouter.POST("", authHandler.Login)
}
