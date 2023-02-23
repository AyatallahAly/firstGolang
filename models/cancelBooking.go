package models

import (
	"log"
)

func (r *Repo) DeleteBooking(app App) {
	//this for only change the cancel col to true instead of deleting all the data
	delsql := `  UPDATE Appointment SET cancel = true 
	WHERE DoctorID = $1 AND PatientID=$2 AND NOT cancel returning * 
	`
	_, delerr := Db.Query(delsql, app.DoctorID, app.PatientID)
	if delerr != nil {
		log.Fatal("Cancel Appointment ", delerr)
	}
}
