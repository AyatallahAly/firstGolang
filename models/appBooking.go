package models

import (
	"log"
)
func (r *Repo) AppBooking(appRead App) {
	q := `INSERT INTO appointment VALUES ($1,$2,$3,$4)`
	_, err := Db.Exec(q,
		appRead.DoctorID,
		appRead.PatientID,
		appRead.StartTime,
		appRead.EndTime,
	)
	if err != nil {
		log.Fatal("Errorrrrrrrrrrrrrrrrrrrrrrrrrr here ", err)
	}
}