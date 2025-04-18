package controllers

import (
	"finalproject/config"
	"finalproject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrderDetailsByOrderID(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var details []models.OrderDetail
	if err := config.DB.Where("order_id = ?", orderID).Find(&details).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, details)
}
