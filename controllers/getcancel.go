package controllers

import (
	"Ass/middlewares"
	"Ass/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCancelBooking(bp models.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {

		//check the roles because only Admin and DR can cancel the app
		userRole, _ := c.Get("role")
		if userRole == middlewares.Patinet {
			c.JSON(http.StatusBadRequest, gin.H{"messang": "You can't cancel your appointment please contact the clinic"})
			return
		}

		bookingBody := newbookRequest{}
		c.ShouldBindJSON(&bookingBody)
		app := models.App{
			DoctorID:  bookingBody.DoctorID,
			PatientID: bookingBody.PatientID,
			StartTime: bookingBody.StartTime,
			EndTime:   bookingBody.EndTime,
		}

		//call delete func
		bp.DeleteBooking(app)

		c.JSON(http.StatusOK, gin.H{"App": "done"})
	}
}
