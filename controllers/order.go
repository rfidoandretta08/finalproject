package controllers

import (
	"finalproject/config"
	"finalproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMyOrders(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var orders []models.Order
	config.DB.Where("customer_id = ?", userID).Preload("OrderDetails").Find(&orders)

	c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Status = "diproses"
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func PayOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.Status = "dikirim"
	config.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{"message": "Payment successful, order is being shipped"})
}

func TrackingOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_id": order.ID,
		"status":   order.Status,
	})
}

func DeliverOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.Status = "selesai"
	config.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{"message": "Order marked as delivered"})
}
