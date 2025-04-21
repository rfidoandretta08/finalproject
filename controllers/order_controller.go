package controllers

import (
	"finalproject/models"
	"finalproject/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService}
}

func (oc *OrderController) GetMyOrders(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	orders, err := oc.orderService.GetUserOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := oc.orderService.CreateOrder(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (oc *OrderController) PayOrder(c *gin.Context) {
	orderID := c.Param("id")

	order, err := oc.orderService.ProcessPayment(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment successful, order is being shipped",
		"order":   order,
	})
}

func (oc *OrderController) TrackingOrder(c *gin.Context) {
	orderID := c.Param("id")

	order, err := oc.orderService.TrackOrder(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_id": order.ID,
		"status":   order.Status,
	})
}

func (oc *OrderController) DeliverOrder(c *gin.Context) {
	orderID := c.Param("id")

	order, err := oc.orderService.CompleteDelivery(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order marked as delivered",
		"order":   order,
	})
}
