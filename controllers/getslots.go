
package controllers

import (
	"Ass/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSlotsBooking(sb models.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingBody := newbookRequest{}
		c.ShouldBindJSON(&bookingBody)
		app := models.App{
			DoctorID:  bookingBody.DoctorID,
			PatientID: bookingBody.PatientID,
			StartTime: bookingBody.StartTime,
			EndTime:   bookingBody.EndTime,
		}
		slotsdata, err := sb.SlotsforeachDR(app)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error on list"})
		}
		c.JSON(http.StatusOK, slotsdata)
	}
}