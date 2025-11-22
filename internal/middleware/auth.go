package middleware

import (
	"log"
	"os"

	"github.com/alebozek/actual-notes/internal/database"
	"github.com/alebozek/actual-notes/internal/models"
	"github.com/gin-gonic/gin"
)

var authLogger = log.New(os.Stdout, "[AUTH] ", log.LstdFlags)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if the cookie is set
		token, err := c.Cookie("token")
		if err != nil {
			authLogger.Println("No valid cookie on host: " + c.ClientIP())
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		// checks if the session is valid
		session := models.Session{Token: token}
		if err := database.DB.Where("token = ? AND expires_at > NOW()", token).First(&session).Error; err != nil {
			authLogger.Println("Invalid session on host: " + c.ClientIP() + ". Token: " + token)
			c.Redirect(302, "/login")
			c.Abort()
			return
		}

		// get the user from the session
		var user models.User
		database.DB.First(&user, session.UserID)

		c.Set("user", user)
		c.Set("user_id", user.ID)
		// goes to the next handler
		c.Next()
	}
}
