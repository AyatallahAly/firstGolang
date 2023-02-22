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
		if userRole != middlewares.Patinet {
			c.JSON(http.StatusBadRequest, gin.H{"Role ": "You are not patient to access this resource"})
			return
		}

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
		if bookingBody.EndTime.Before(bookingBody.StartTime) {
			c.JSON(http.StatusOK, gin.H{"message": "End Time has to be after start time "})
			return
		}

		//check the app not less than 15 min and not more than 2 hours
		askingTime := bookingBody.EndTime.Sub(bookingBody.StartTime)
		if (askingTime.Minutes() < 15) || (askingTime.Hours() > 2) {
			c.JSON(http.StatusOK, gin.H{"Appoinment not Booked Minimum Duration: 15 mins, Max Duration: 120 mins": askingTime.Minutes()})
			return
		}

		//get count of app for dr on same day
		var countDrApp int
		countsql := `select count(date(startTime))from appointment WHERE DoctorID=$1 AND date(startTime)=$2`
		counterr := models.Db.QueryRow(countsql, app.DoctorID, app.StartTime).Scan(&countDrApp)
		if counterr != nil {
			c.JSON(http.StatusOK, gin.H{"error in query Doctor time reserved": "check the DB "})
		}

		// if no of app ==0 add the new app
		if countDrApp == 0 {
			bp.AppBooking(app)
			c.JSON(http.StatusOK, gin.H{"Appointement for Doctor ": bookingBody.DoctorID})
			return
		}

		//check if this patient has more than one app at same day, patient can't book one more
		var countPatientID int
		countPatientSql := `select count (starttime) from appointment where DoctorID =$1
		AND patientid=$2  
		AND date(startTime)=$3`
		countPatientErr := models.Db.QueryRow(countPatientSql, app.DoctorID, app.PatientID, app.StartTime).Scan(&countPatientID)
		if countPatientErr != nil {
			c.JSON(http.StatusOK, gin.H{"error in query Patient ": "This Patient has an appointment with this Dr on this Day "})
		}

		if countPatientID != 0 {
			c.JSON(http.StatusOK, gin.H{"message": "This patient already had an appointement on this day"})
			return
		}

		//sum of working hours per Dr
		var counthour float64
		counthourSql := `select sum(EXTRACT(EPOCH FROM (endtime - starttime))/3600)from Appointment where Doctorid=$1 AND date(startTime)=$2`
		counthourerr := models.Db.QueryRow(counthourSql, app.DoctorID, app.StartTime).Scan(&counthour)
		if counthourerr != nil {
			c.JSON(http.StatusOK, gin.H{"error in query  total time fpor ecah Dr": "check the DB "})
		}

		//check if sum of working hours more than 8h or no of apps more than 12 apps
		if countDrApp >= 12 || counthour >= 8 {
			c.JSON(http.StatusOK, gin.H{"Doctor is fully Booked this day ": bookingBody.DoctorID})
			return
		}

		//check overlapping before reserve new app
		var overlap int
		overLapsql := `select count (starttime) from appointment where DoctorID=$1 AND(SELECT($2,$3) 
		overlaps (startTime, endTime)is true)`
		overLaperr := models.Db.QueryRow(overLapsql, app.DoctorID, app.StartTime, app.EndTime).Scan(&overlap)
		if overLaperr != nil {
			c.JSON(http.StatusOK, gin.H{"message": "overlap time error"})
			return
		}

		if overlap == 0 {
			bp.AppBooking(app)
			c.JSON(http.StatusOK, gin.H{"Appointement for Doctor ": bookingBody.DoctorID})
		} else {
			c.JSON(http.StatusOK, gin.H{"Doctor is booked on this time ": "Please choose another time "})
		}
	}
}
