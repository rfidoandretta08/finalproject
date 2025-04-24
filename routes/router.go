package routes

import (
	"finalproject/controllers"
	"finalproject/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	customerController *controllers.CustomerController,
	productController *controllers.ProductController,
	orderController *controllers.OrderController,
	userController *controllers.UserController,
	categoryController *controllers.CategoryController,
	reportingController *controllers.ReportingController,
) {
	// Auth (public)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/catalog", productController.GetProducts)

	// === CUSTOMER ROUTES ===
	customer := r.Group("/")
	customer.Use(middleware.AuthMiddleware("customer"))
	{
		// CUSTOMER
		customer.GET("/customers/:id", customerController.GetCustomerByID)
		customer.POST("/customers", customerController.CreateCustomer)
		customer.PUT("/customers/:id", customerController.UpdateCustomer)

		// PRODUCT
		customer.GET("/products", productController.GetProducts)

		// ORDER
		customer.GET("/orders/user/:userID", orderController.GetUserOrders)
		customer.GET("/orders/:id", orderController.TrackOrder)
		customer.POST("/orders", orderController.CreateOrder)
		customer.PUT("/orders/:id/payment", orderController.ProcessPayment)

		// USER
		customer.GET("/users/:id", userController.GetUserByID)
		customer.PUT("/users/:id", userController.UpdateUser)

		// CATEGORY
		customer.GET("/categories", categoryController.GetAllCategories)

	}

	// === ADMIN ROUTES ===
	admin := r.Group("/")
	admin.Use(middleware.AuthMiddleware("admin"))
	{
		// CATEGORY
		admin.POST("/categories", categoryController.CreateCategory)
		admin.GET("/categories/all", categoryController.GetAllCategories)
		admin.DELETE("/categories/:id", categoryController.DeleteCategory)

		// CUSTOMER
		admin.GET("/customers", customerController.GetAllCustomers)
		admin.DELETE("/customers/:id", customerController.DeleteCustomer)

		// PRODUCT
		admin.GET("/products/all", productController.GetProducts)
		admin.POST("/products", productController.CreateProduct)
		admin.PUT("/products/:id", productController.UpdateProduct)
		admin.DELETE("/products/:id", productController.DeleteProduct)

		// ORDER
		admin.PUT("/orders/:id/complete", orderController.CompleteDelivery)
		admin.GET("/orders/order/:userID", orderController.GetUserOrders)

		// USER
		admin.GET("/users", userController.GetAllUsers)
		admin.DELETE("/users/:id", userController.DeleteUser)

		// REPORTING
		admin.GET("/reports/products-by-quantity", reportingController.AllProductsByQuantity)
		admin.GET("/reports/customer-spendings", reportingController.CustomerSpendings)
		admin.GET("/reports/products-by-nominal", reportingController.ProductsByNominal)
	}

}
