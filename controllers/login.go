package controllers

import (
	"Ass/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginGet(l models.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Get email & password
		loginBody := newUserRequest{}
		if c.Bind(&loginBody) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Falid to read Body",
			})
			return
		}

		user := models.User{
			//Username: loginBody.Username,
			Email:    loginBody.Email,
			Password: loginBody.Password,
			Role:     loginBody.Role,
		}

		// call db and look up the user
		dataq := `select * from person where email = $1`
		err := models.Db.QueryRow(dataq, loginBody.Email).Scan(
			&loginBody.ID,
			&loginBody.Email,
			&loginBody.Password,
			&loginBody.Role,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "call db and get user"})
		}

		//compare sent in pass with saved user pass hash
		passerr := bcrypt.CompareHashAndPassword([]byte(loginBody.Password), []byte(user.Password))
		if passerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or Password_2"})
			return
		}

		// Create a new token object, specifying signing method and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":  user.Email,
			"role": user.Role,
			"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, tokenerr := token.SignedString([]byte(os.Getenv("SECRET")))
		if tokenerr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token erorr"})
		}
		// c.JSON(http.StatusOK, gin.H{
		// 	"token": tokenString,
		// })

		l.LoginUser(user)
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

		// c.Status(http.StatusNoContent)
		//c.JSON(http.StatusOK,gin.H{"user has been found his role is ": user.Role})
	}
}
