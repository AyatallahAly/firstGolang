package models

import (
	"log"
	"time"
)

func (r *Repo) SlotsforeachDR(app App) ([]Slot, error) {
	var drstTime, drETime time.Time

	DrqTime := `
	select DrStartTime,DrEndTime from DrTimes 
	WHERE TimesDID=$1 AND date(DrStartTime)=$2;`
	DrqTimeErr := Db.QueryRow(DrqTime, app.DoctorID, app.StartTime).Scan(&drstTime, &drETime)
	if DrqTimeErr != nil {
		log.Fatal("Doctor start and end time for Appointment")
	}

	q := `
	select startTime,endtime  from appointment 
	WHERE DoctorID=$1 AND date(startTime)=$2
	ORDER BY StartTime;`
	rows, err := Db.Query(q, app.DoctorID, app.StartTime)
	if err != nil {
		log.Fatal("unable to execute search query", err)
		return nil, err
	}
	//defer rows.Close()
	var slots []Slot
	// var i =0
	for rows.Next() {
		var t Slot
		// Handle result set.
		if err := rows.Scan(&t.StartTime, &t.EndTime); err != nil {
			log.Fatal("unable to scan time", err)
			return slots, err
		}
		slots = append(slots, t)
	}

	nOfArr := len(slots) + 1
	avlTime := make([]Slot, nOfArr)
	for i := 0; i < nOfArr; i++ {
		if i == 0 {
			a := drstTime.Add(slots[0].StartTime.Sub(drstTime))
			avlTime[0].StartTime = drstTime
			avlTime[0].EndTime = a
		} else if i == nOfArr-1 {
			b := drETime.Add(slots[i-1].EndTime.Sub(drETime))
			avlTime[i].StartTime = b
			avlTime[i].EndTime = drETime

		} else {
			avlTime[i].StartTime = slots[i-1].EndTime
			avlTime[i].EndTime = slots[i].StartTime

		}

	}

	// Check for any error in either result set.
	if err := rows.Err(); err != nil {
		log.Fatal("Rows error", err)
		return slots, err
	}
	return avlTime[:], err
}