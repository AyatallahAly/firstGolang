
package controllers

import (
	"Ass/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetCancelBooking(bp models.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingBody := newbookRequest{}
		c.ShouldBindJSON(&bookingBody)
		app := models.App{
			DoctorID:  bookingBody.DoctorID,
			PatientID: bookingBody.PatientID,
			StartTime: bookingBody.StartTime,
			EndTime:   bookingBody.EndTime,
		}
		bp.DeleteBooking(app)
		c.JSON(http.StatusOK, gin.H{"App": "done"})
	}
}
