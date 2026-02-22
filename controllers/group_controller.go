package controllers

import (
	"expense-tracker-go/config"
	"expense-tracker-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGroup(c *gin.Context) {

	var group models.Group

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&group)

	c.JSON(http.StatusCreated, group)
}
