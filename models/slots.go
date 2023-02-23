package models

import (
	"log"
	"time"
)

func (r *Repo) SlotsforeachDR(app App) ([]Slot, error) {
	var drstTime, drETime time.Time
    
	//get the working hours for dr for speacific date 
	DrqTime := `
	select DrStartTime,DrEndTime from DrTimes 
	WHERE DoctorID=$1 AND date(DrStartTime)=$2`

	DrqTimeErr := Db.QueryRow(DrqTime, app.DoctorID, app.StartTime).Scan(&drstTime, &drETime)
	if DrqTimeErr != nil {
		log.Fatal("error in start & end time query")
	}
	
	// get starttime and endtime for specific DR and for specifoc date 
	q := `
	select startTime,endtime from appointment 
	WHERE DoctorID= $1 AND date(startTime)=$2 AND cancel ='false'
	ORDER BY StartTime`
	
	rows, err := Db.Query(q, app.DoctorID, app.StartTime)
	if err != nil {
		log.Fatal("unable to execute search query", err)
		return nil, err
	}
	//defer rows.Close()
	
	var slots []Slot
	for rows.Next() {
		var t Slot
		// Handle result set.
		if err := rows.Scan(&t.StartTime, &t.EndTime); err != nil {
			log.Fatal("unable to scan time", err)
			return slots, err
		}
		slots = append(slots, t)
	}

	// cal the gab between each 2 app.
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

	//save only avilable time for Dr
	finalavlTime := make([]Slot, nOfArr)
	for j:= 0; j < nOfArr ; j++{
		if (avlTime[j].StartTime!=avlTime[j].EndTime){
			finalavlTime[j].StartTime= avlTime[j].StartTime
			finalavlTime[j].EndTime =avlTime[j].EndTime
		}
	}

	// Check for any error in either result set.	
	if err := rows.Err(); err != nil {
		log.Fatal("Rows error", err)
		return slots, err
	}
	
	return finalavlTime[:], err
}