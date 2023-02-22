package controllers

//package controllers

import (
	//"Ass/middlewares"
	"Ass/models"
	"net/http"
	//"time"
	"github.com/gin-gonic/gin"
)


func GetDrById(dr models.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		doctorBody := newDoctorData{}
		c.ShouldBindJSON(&doctorBody)

		doctor := models.Doctor{
			Docid:      doctorBody.Docid,
			Person_id:  doctorBody.Person_id,
			Weekday:    doctorBody.Weekday,
			Speciality: doctorBody.Speciality,
		}

		// call db and get user
		drq := `select PersonID,weekday,Speciality from doctor where DID= $1`
		err0 := models.Db.QueryRow(drq, doctorBody.Docid).Scan(
			&doctorBody.Person_id,
			&doctorBody.Weekday,
			&doctorBody.Speciality,
		)
		if err0 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "call db and get doctordata"})
		}

		//return data and print it out
		finaldata, err := dr.DoctorsdataById(doctor.Docid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error on list"})
		}

		c.JSON(http.StatusOK, finaldata)

	}
}
