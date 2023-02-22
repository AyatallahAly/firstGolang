package models

import (
	"log"
)

func (r *Repo) NoOfBookingPerDr(app App) (countDrApp int) {
	//check all reserved time for each dr
	//var countDrApp int
	countsql := `select count(date(startTime))from appointment WHERE DoctorID=$1 AND date(startTime)=$2`
	counterr := Db.QueryRow(countsql, app.DoctorID, app.StartTime).Scan(&countDrApp)
	if counterr != nil {
		log.Fatal("Doctor has the max number for Appointment", countDrApp)
	}
	return countDrApp
}