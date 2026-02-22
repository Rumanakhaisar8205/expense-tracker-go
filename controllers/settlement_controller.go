package controllers

import (
	"expense-tracker-go/algorithm"
	"expense-tracker-go/config"
	"expense-tracker-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSettlement(c *gin.Context) {

	groupID := c.Param("id")

	var expenses []models.Expense
	config.DB.Where("group_id = ?", groupID).Find(&expenses)

	balanceMap := make(map[uint]float64)

	for _, e := range expenses {
		balanceMap[e.PaidBy] += e.Amount
	}

	var balances []algorithm.Balance

	for user, amount := range balanceMap {
		balances = append(balances, algorithm.Balance{
			UserID: user,
			Amount: amount,
		})
	}

	result := algorithm.Settle(balances)

	c.JSON(http.StatusOK, result)
}
