package controllers

import (
	"Ass/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListDoc(dr models.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {

		doctorsdata, err := dr.ListDoc()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error on list"})
		}

		c.JSON(http.StatusOK, doctorsdata)

	}
}
