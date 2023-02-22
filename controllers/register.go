package controllers

import (
	"Ass/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func RegisterPost(p models.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newUserRequest{}
		if c.Bind(&requestBody) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Falid to read Body",
			})
			return
		}

		user := models.User{
			//Username: requestBody.Username,
			Email:    requestBody.Email,
			Password: requestBody.Password,
			Role:     requestBody.Role,
		}

		//hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password encryp"})
		}
		//create user with hash password
		user.Password = string(hashedPassword)

		//p.BeforeSave(user)
		p.RegisterUser(user)

		//respond
		c.Status(http.StatusNoContent)

	}
}
