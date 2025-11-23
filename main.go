package main

import (
	"os"

	"github.com/alebozek/actual-notes/internal/database"
	"github.com/alebozek/actual-notes/internal/handlers"
	"github.com/alebozek/actual-notes/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	handlers.DOMAIN = os.Getenv("DOMAIN")

	database.InitDB()
	router := gin.Default()
	router.Use(gin.Recovery())
	router.StaticFile("/favicon.ico", "static/favicon.ico")
	router.Static("/static", "static/")
	router.LoadHTMLGlob("internal/templates/*")

	router.GET("/login", handlers.LoginPage())
	router.POST("/login", handlers.Login())

	router.GET("/register", handlers.RegisterPage())
	router.POST("/register", handlers.Register())

	// authorized pages
	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// here will be the protected endpoints
	}

	router.Run(":8080")
}
