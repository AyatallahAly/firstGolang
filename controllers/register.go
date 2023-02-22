package controllers

import (
	"Ass/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterPost(p models.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		// read the user data (Email,Password,role) 
		requestBody := newUserRequest{}
		if c.Bind(&requestBody) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Falid to read Body",
			})
			return
		}

		user := models.User{
			Email:    requestBody.Email,
			Password: requestBody.Password,
			Role:     requestBody.Role,
		}
		
		// save the hashed password  
		user.Password = p.BeforeSave(user)
		
		// call the registeratin fun from models 
		p.RegisterUser(user)

		//respond
		c.Status(http.StatusNoContent)
	}
}
