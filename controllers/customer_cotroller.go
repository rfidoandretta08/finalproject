package controllers

import (
	"finalproject/models"

	"finalproject/services"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	service service.CustomerService
}

func NewCustomerController(service service.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

func (ctrl *CustomerController) GetAllCustomers(c *gin.Context) {
	customers, err := ctrl.service.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (ctrl *CustomerController) CreateCustomer(c *gin.Context) {
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateCustomer(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (ctrl *CustomerController) UpdateCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer, err := ctrl.service.GetCustomerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if err := c.ShouldBindJSON(customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.UpdateCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (ctrl *CustomerController) DeleteCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := ctrl.service.DeleteCustomer(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
