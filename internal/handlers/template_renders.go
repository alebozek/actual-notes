package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{"title": "Register"})
	}
}
