package controllers

import (
	"finalproject/models"
	"finalproject/services"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService}
}

// GET /orders/user/:userID
func (h *OrderController) GetUserOrders(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID tidak valid"})
		return
	}

	orders, err := h.orderService.GetUserOrders(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// POST /orders
func (h *OrderController) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hitung total & subtotal
	if err := h.orderService.CalculateHargaDanTotal(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kurangi stok
	if err := h.orderService.KurangiStok(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan order
	if err := h.orderService.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// PUT /orders/:id/payment
func (h *OrderController) ProcessPayment(c *gin.Context) {
	orderID := c.Param("id")
	order, err := h.orderService.ProcessPayment(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

// GET /orders/:id
func (h *OrderController) TrackOrder(c *gin.Context) {
	orderID := c.Param("id")
	order, err := h.orderService.TrackOrder(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderController) CompleteDelivery(c *gin.Context) {
	orderID := c.Param("id")
	order, err := h.orderService.CompleteDelivery(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}
