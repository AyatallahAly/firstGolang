package models

import (
	"log"
)

func (r *Repo) DeleteBooking(app App) {
	delsql := `delete from Appointment where DoctorID=$1 AND PatientID=$2 returning *`
	_, delerr := Db.Query(delsql, app.DoctorID, app.PatientID)
	if delerr != nil {
		log.Fatal("Cancel Appointment ", delerr)
	}
}
