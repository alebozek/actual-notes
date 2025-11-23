package handlers

import (
	"net/http"
	"time"

	"github.com/alebozek/actual-notes/internal/database"
	"github.com/alebozek/actual-notes/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var DOMAIN string

// Handler for register view
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get form results
		username := c.PostForm("username")
		password := c.PostForm("password")

		var checkUser models.User
		userCheckRes := database.DB.Where("username = ?", username).Find(&checkUser)
		if userCheckRes.RowsAffected > 0 {
			c.HTML(302, "register.html", gin.H{"error": "User already exists"})
			return
		}

		// generate hash
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.HTML(http.StatusOK, "register.html", gin.H{"error": err.Error()})
			return
		}

		user := models.User{Username: username, Password: string(hash), CreatedAt: time.Now()}

		res := database.DB.Create(&user)
		if res.Error != nil {
			c.HTML(http.StatusOK, "register.html", gin.H{"error": res.Error.Error()})
			return
		}

		c.Redirect(302, "/login")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// get hash
		var userCheck models.User
		res := database.DB.Where("username = ?", username).Find(&userCheck)
		if res.RowsAffected != 1 {
			c.HTML(http.StatusOK, "login.html", gin.H{"error": "Invalid credentials. Try again."})
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(userCheck.Password), []byte(password))
		if err != nil {
			c.HTML(302, "login.html", gin.H{"error": "Invalid credentials. Try again."})
			return
		} else {
			// generate UUID for token
			token := uuid.New().String()

			// create session
			session := models.Session{Token: token, UserID: userCheck.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()}

			// save session
			sessionRes := database.DB.Create(&session)
			if sessionRes.RowsAffected != 1 {
				c.HTML(302, "login.html", gin.H{"error": "Couldn't create session. Please try again."})
				return
			} else {
				c.SetCookie("token", token, 604800, "/", DOMAIN, true, true)
				c.HTML(http.StatusTemporaryRedirect, "/dashboard", gin.H{})
				return
			}
		}

	}
}
