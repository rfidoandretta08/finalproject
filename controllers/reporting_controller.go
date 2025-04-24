package controllers

import (
	"net/http"

	"finalproject/services"

	"github.com/gin-gonic/gin"
)

type ReportingController struct {
	svc services.ReportingService
}

func NewReportingController(svc services.ReportingService) *ReportingController {
	return &ReportingController{svc: svc}
}

func (c *ReportingController) Register(r *gin.Engine) {
	rpt := r.Group("/reports")
	{
		rpt.GET("/products-by-quantity", c.AllProductsByQuantity)
		rpt.GET("/customer-spendings", c.CustomerSpendings)
		rpt.GET("/products-by-nominal", c.ProductsByNominal)
	}
}

func (c *ReportingController) AllProductsByQuantity(ctx *gin.Context) {
	data, err := c.svc.ProductsByQuantity()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *ReportingController) CustomerSpendings(ctx *gin.Context) {
	data, err := c.svc.GetCustomerSpendings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *ReportingController) ProductsByNominal(ctx *gin.Context) {
	data, err := c.svc.ProductsByNominal()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
