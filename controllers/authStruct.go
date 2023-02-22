package controllers

//package controllers

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

//	"fmt"

//"log"
//	"net/http"

//"github.com/gin-gonic/gin"

//type NewUser interface{
//	GetUserName() string
//}

//func(u *User) GetUserName()string{
//	return u.name
//}

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




// func CheckUserRole(c *gin.Context, role string) {
// 	userRole  := c.GetString("user_role")
//     if userRole != role {
//     c.JSON(http.StatusBadRequest,gin.H{"Role Error":"You are not authorised to access this resource"})
//     }  
// }

func Validate(c *gin.Context){
	// var role []string 
	// userRole, _ := c.Get("role")
    // if userRole == middlewares.Admin {
	// 	c.JSON(http.StatusBadRequest,gin.H{"Role ":"You are Admin authorised to access this resource"})
		
	// 	return
    // }

	// if userRole == middlewares.Doctor {
	// 	c.JSON(http.StatusBadRequest,gin.H{"Role Error":"You are Doctor authorised to access this resource"})
	// 	return
	// }

	// if userRole == middlewares.Patinet {
	// 	c.JSON(http.StatusBadRequest,gin.H{"Role Error":"You are Patient authorised to access this resource"})
	// 	return
	// }

	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user, 
	})
}




