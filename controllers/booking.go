package controllers

import (
	"Ass/middlewares"
	"Ass/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



func GetAppBooking(bp models.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, _ := c.Get("role")
		if userRole == middlewares.Patinet {
		
		bookingBody := newbookRequest{}
		c.ShouldBindJSON(&bookingBody)

		app := models.App{
			//AppId:     bookingBody.AppId,
			DoctorID:  bookingBody.DoctorID,
			PatientID: bookingBody.PatientID,
			StartTime: bookingBody.StartTime,
			EndTime:   bookingBody.EndTime,
		}
		//Check end time after start time
		// if start time before end time
		// c.JSON(http.StatusOK, gin.H{"message": "End Time has to be after start time "})
		// return
		// askingTime := bookingBody.EndTime.Sub(bookingBody.StartTime)
		

		if bookingBody.EndTime.After(bookingBody.StartTime) {
			askingTime := bookingBody.EndTime.Sub(bookingBody.StartTime)
			//Appointment time betweeen not less than 15 mins and not more than  2 hours
			if (askingTime.Minutes() >= 15) && (askingTime.Hours() <= 2) {
				//check all reserved time for each Dr
				var countDrApp int
				countsql := `select count(date(startTime))from appointment WHERE DoctorID=$1 AND date(startTime)=$2`
				counterr := models.Db.QueryRow(countsql, app.DoctorID, app.StartTime).Scan(&countDrApp)
				if counterr != nil {
					c.JSON(http.StatusOK, gin.H{"error in query Doctor time reserved": "check the DB "})
				}
				var countPatientID int
				countPatientSql := `select count (starttime) from appointment where DoctorID =$1
				 AND patientid=$2  
				 AND date(startTime)=$3`
				countPatientErr := models.Db.QueryRow(countPatientSql, app.DoctorID, app.PatientID, app.StartTime).Scan(&countPatientID)
				if countPatientErr != nil {
					c.JSON(http.StatusOK, gin.H{"error in query Patient ": "This Patient has an appointment with this Dr on this Day "})
				 }
				var counthour int
				counthourSql := `select sum(EXTRACT(EPOCH FROM (endtime - starttime))/3600) 
				as diff_time_hours from Appointment where Doctorid=$1 AND date(startTime)=$2`
				counthourerr := models.Db.QueryRow(counthourSql, app.DoctorID, app.StartTime).Scan(&counthour)
				if counthourerr != nil {
					c.JSON(http.StatusOK, gin.H{"error in query  total time fpor ecah Dr": "check the DB "})
				}
				if countDrApp <= 12 && counthour <= 8 && countPatientID == 0 {
					var overlap int
					overLapsql := `select count (starttime) from appointment where DoctorID=$1 AND(SELECT($2,$3) 
					overlaps (startTime, endTime)is true)`
					overLaperr := models.Db.QueryRow(overLapsql, app.DoctorID, app.StartTime, app.EndTime).Scan(&overlap)
					if overLaperr != nil {
						c.JSON(http.StatusOK, gin.H{"Checking Overlaping Time ": "wait "})
					}
					if overlap == 0 {
						bp.AppBooking(app)
						c.JSON(http.StatusOK, gin.H{"Appointement for Doctor ": bookingBody.DoctorID})

					} else {
						c.JSON(http.StatusOK, gin.H{"Doctor is booked on this time ": "Please choose another time "})

					}
				} else {
					c.JSON(http.StatusOK, gin.H{"Doctor is fully Booked this day ": bookingBody.DoctorID})
				}
			} else {
				c.JSON(http.StatusOK, gin.H{"Appoinment not Booked Minimum Duration: 15 mins, Max Duration: 120 mins": askingTime.Minutes()})

			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "End Time has to be after start time "})
		}
	} else{
		c.JSON(http.StatusBadRequest,gin.H{"Role ":"You are not patient to access this resource"})
		return
	}
	}
}