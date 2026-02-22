package main

import (
	"expense-tracker-go/config"
	"expense-tracker-go/controllers"
	"expense-tracker-go/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// DB Connect
	config.ConnectDB()

	// Auto Migrate
	config.DB.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.Expense{},
		&models.ExpenseShare{}, // ADD THIS
	)

	r := gin.Default()

	// User Routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)

	// Group Routes
	r.POST("/groups", controllers.CreateGroup)

	// Expense Routes
	r.POST("/expenses", controllers.AddExpense)
	r.GET("/groups/:id/expenses", controllers.GetGroupExpenses)

	r.GET("/groups/:id/settle", controllers.GetSettlement)

	r.Run(":8080")
}
