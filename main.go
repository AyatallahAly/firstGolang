package main

import (
	"Ass/controllers"
	"Ass/middlewares"
	//"Ass/middlewares"

	//"Ass/middlewares"
	"Ass/models"

	//"fmt"
	//"time"

	//"net/http"

	//"log"
	//"os/user"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/lib/pq"
)

func main() {

	models.ConnectDataBase()
	testuser := models.New()
	r := gin.Default()
	{
		public := r.Group("/api")

		//public.POST("/register", controllers.RegisterPost(testuser))
		public.POST("/login", controllers.LoginGet(testuser))
		//public.GET("/validate",middlewares.RequiredAuth, controllers.Validate)
		//public.GET("/doclist", controllers.GetListDoc(testuser))
		//public.GET("/docdataById", controllers.GetDrById(testuser))
		public.GET("/booking",middlewares.RequiredAuth ,controllers.GetAppBooking(testuser))
		//public.GET("/cancelBooking", controllers.GetCancelBooking(testuser))
		//public.GET("/doctors/:doctorId/slots", controllers.GetSlotsBooking(testuser))
		//protected.Use(middlewares.JwtAuthMiddleware())
		//protected.GET("/user",controllers.CurrentUser)
		defer models.Db.Close()

	}

	r.Run(":8080")

}
