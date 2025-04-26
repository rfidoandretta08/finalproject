package main

import (
	"finalproject/config"
	"finalproject/controllers"
	"finalproject/repositories"
	"finalproject/routes"
	"finalproject/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi database
	config.InitDB()

	// Inisialisasi Gin
	r := gin.Default()

	// Inisialisasi repository
	customerRepository := repositories.NewCustomerRepository(config.DB)
	productRepository := repositories.NewProductRepository(config.DB)
	orderRepository := repositories.NewOrderRepository(config.DB)
	NewUserRepository := repositories.NewUserRepository(config.DB)
	reportingRepository := repositories.NewReportingRepository(config.DB)

	// Inisialisasi service dengan repository
	customerService := services.NewCustomerService(customerRepository)
	productService := services.NewProductService(productRepository)
	orderService := services.NewOrderService(orderRepository, productRepository)
	userService := services.NewUserService(NewUserRepository)
	reportingService := services.NewReportingService(reportingRepository)

	customerController := controllers.NewCustomerController(customerService)
	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService)
	userController := controllers.NewUserController(userService)
	categoryController := controllers.NewCategoryController()
	reportingController := controllers.NewReportingController(reportingService)

	// Setup semua routes
	routes.SetupRoutes(r, customerController, productController, orderController, userController, categoryController, reportingController)
	// Run server
	r.Run(":8080")
}
