package models

import (
	"log"
)

func (r *Repo) DoctorsdataById(id int) ([]Doctor, error) {
	var d Doctor
	d.Docid = id
	drq := `select PersonID,weekday,Speciality from doctor where DID =$1`
	err := Db.QueryRow(drq, id).Scan(&d.Person_id, &d.Weekday, &d.Speciality)
	if err != nil {
		log.Fatal("unable to execute search query", err)
		return nil, err
	}
	var doctors []Doctor
	doctors = append(doctors, d)
	return doctors, err
}
