package controllers

import (
	"time"
)

type newUserRequest struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type newDoctorData struct {
	Docid      int    `json:"Docid"`
	Person_id  int    `json:"person_id"`
	Weekday    string `json:"weekday"`
	Speciality string `json:"speciality"`
}

type newbookRequest struct {
	DoctorID  int64     `json:"DoctorID"`
	PatientID int64     `json:"PatientID"`
	StartTime time.Time `json:"StartTime"`
	EndTime   time.Time `json:"EndTime"`
	Cancel    bool
}
