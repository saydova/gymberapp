package main

import (
	"gymberapp/controllers"
	"gymberapp/middleware"
	"gymberapp/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	dsn := "root:password@tcp(127.0.0.1:3306)/gymdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.Member{}, &models.Trainer{}, &models.Class{}, &models.Enrollment{}, &models.Payment{})

	router := gin.Default()

	router.POST("/login", controllers.Login)

	api := router.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("/members", controllers.GetMembers)
		api.POST("/members", controllers.CreateMember)
		api.PUT("/members/:id", controllers.UpdateMember)
		api.DELETE("/members/:id", controllers.DeleteMember)

		api.GET("/trainers", controllers.GetTrainers)
		api.POST("/trainers", controllers.CreateTrainer)
		api.PUT("/trainers/:id", controllers.UpdateTrainer)
		api.DELETE("/trainers/:id", controllers.DeleteTrainer)

		api.GET("/classes", controllers.GetClasses)
		api.POST("/classes", controllers.CreateClass)
		api.PUT("/classes/:id", controllers.UpdateClass)
		api.DELETE("/classes/:id", controllers.DeleteClass)

		api.GET("/enrollments", controllers.GetEnrollments)
		api.POST("/enrollments", controllers.CreateEnrollment)
		api.PUT("/enrollments/:id", controllers.UpdateEnrollment)
		api.DELETE("/enrollments/:id", controllers.DeleteEnrollment)

		api.GET("/payments", controllers.GetPayments)
		api.POST("/payments", controllers.CreatePayment)
		api.PUT("/payments/:id", controllers.UpdatePayment)
		api.DELETE("/payments/:id", controllers.DeletePayment)

		router.POST("/login", controllers.Login)
		router.POST("/register", controllers.Login)
	}

	router.Run(":8080")
}
