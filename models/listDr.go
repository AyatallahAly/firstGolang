package models

import (
	"log"
)
func (r *Repo) ListDoc() ([]Doctor, error) {
	rows, err := Db.Query("SELECT * from doctor")
	if err != nil {
		log.Fatal("unable to execute search query", err)
		return nil, err
	}
	defer rows.Close()
	///// Doctor to hold data from returned rows
	var doctors []Doctor
	// Loop through the first result set.
	for rows.Next() {
		var d Doctor
		// Handle result set.
		if err := rows.Scan(&d.Docid, &d.Person_id, &d.Weekday, &d.Speciality); err != nil {
			return doctors, err
		}
		//fmt.Print(d.Docid, d.person_id, d.weekday)
		doctors = append(doctors, d)

	}
	// Check for any error in either result set.
	if err := rows.Err(); err != nil {
		return doctors, err
	}

	return doctors, err

}