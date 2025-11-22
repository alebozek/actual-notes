package handlers

import (
	"net/http"
	"time"

	"github.com/alebozek/actual-notes/internal/database"
	"github.com/alebozek/actual-notes/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Handler for register view
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get form results
		username := c.PostForm("username")
		password := c.PostForm("password")

		var checkUser models.User
		userCheckRes := database.DB.Where("username = ?", username).Find(&checkUser)
		if userCheckRes.RowsAffected > 0 {
			c.HTML(302, "register.html", gin.H{"title": "Register", "error": "User already exists"})
			return
		}

		// generate hash
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.HTML(http.StatusOK, "register.html", gin.H{"title": "Register", "error": err.Error()})
			return
		}

		user := models.User{Username: username, Password: string(hash), CreatedAt: time.Now()}

		res := database.DB.Create(&user)
		if res.Error != nil {
			c.HTML(http.StatusOK, "register.html", gin.H{"title": "Register", "error": res.Error.Error()})
			return
		}
		database.DBLogger.Println(res.RowsAffected)

		c.Redirect(302, "/login")
	}
}
