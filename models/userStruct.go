package models

import (
	"time"
)

type User struct {
	ID       int
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
	Role  string `json:"role" db:"role" binding:"required"`
}

type Doctor struct {
	Docid      int    `json:"Docid" db:"Docid" binding:"required"`
	Person_id  int    `json:"person_id" db:"person_id" binding:"required" `
	Weekday    string `json:"weekday" db:"weekday" binding:"required"`
	Speciality string `json:"speciality" db:"Speciality" binding:"required"`
}

type App struct {
	DoctorID  int64     `json:"DoctorID" db:"DoctorID"  binding:"required"`
	PatientID int64     `json:"PatientID" db:"PatientID"  binding:"required"`
	StartTime time.Time `json:"StartTime" db:"StartTime" binding:"required"`
	EndTime   time.Time `json:"EndTime" db:"EndTime" binding:"required"`
	Cancel    bool
}

type Slot struct {
	StartTime time.Time `json:"StartTime" db:"StartTime" binding:"required"`
	EndTime   time.Time `json:"EndTime" db:"EndTime" binding:"required"`
}

type Repo struct {
	Users   []User
	Docrors []Doctor
	Apps    []App
	Slots   []Slot
}
type Adder interface {
	//BeforeSave(user User)
	RegisterUser(user User)
	LoginUser(user User)
	ListDoc() ([]Doctor, error)
	DoctorsdataById(id int) ([]Doctor, error)
	AppBooking(appRead App)
	NoOfBookingPerDr(app App) (countDrApp int)
	DeleteBooking(app App)
	SlotsforeachDR(app App) ([]Slot, error)
}

func New() *Repo {
	return &Repo{
		Users:   []User{},
		Docrors: []Doctor{},
		Apps:    []App{},
		Slots:   []Slot{},

	}
}

// func (r *Repo) BeforeSave(user User) error {

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Fatal("unable to hash password ", err)
// 		return err
// 	}
// 	user.Password = string(hashedPassword)
// 	return nil
// }






