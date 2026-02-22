package controllers

import (
	"expense-tracker-go/config"
	"expense-tracker-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {

	var expense models.Expense

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&expense)

	c.JSON(http.StatusCreated, expense)
}

func GetGroupExpenses(c *gin.Context) {

	var expenses []models.Expense

	groupID := c.Param("id")

	config.DB.Where("group_id = ?", groupID).Find(&expenses)

	c.JSON(200, expenses)
}
