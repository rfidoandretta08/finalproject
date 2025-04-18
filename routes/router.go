package routes

import (
	"finalproject/controllers" // Pastikan untuk mengimpor controller yang benar
	"finalproject/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public Auth
	r.POST("/register", controllers.Register) // Memastikan Register sudah terdefinisi di controller
	r.POST("/login", controllers.Login)       // Memastikan Login sudah terdefinisi di controller

	// Customer
	customer := r.Group("/orders")
	{
		customer.Use(middleware.AuthMiddleware("customer"))
		customer.GET("/my", controllers.GetMyOrders)             // Pastikan GetMyOrders ada di controller order
		customer.POST("/customer", controllers.CreateOrder)      // Pastikan CreateOrder ada di controller order
		customer.PUT("/:id/pay", controllers.PayOrder)           // Pastikan PayOrder ada di controller order
		customer.GET("/:id/tracking", controllers.TrackingOrder) // Pastikan TrackingOrder ada di controller order
	}

	// Product (public)
	product := r.Group("/products")
	{
		product.GET("/", controllers.GetProducts) // Pastikan GetProducts ada di controller product
	}

	// Admin
	admin := r.Group("/admin")
	{
		admin.Use(middleware.AuthMiddleware("admin"))

		// Admin Products
		adminProducts := admin.Group("/products")
		{
			adminProducts.POST("/", controllers.CreateProduct)      // Pastikan CreateProduct ada di controller product
			adminProducts.GET("/", controllers.GetProducts)         // Pastikan GetProducts ada di controller product
			adminProducts.PUT("/:id", controllers.UpdateProduct)    // Pastikan UpdateProduct ada di controller product
			adminProducts.DELETE("/:id", controllers.DeleteProduct) // Pastikan DeleteProduct ada di controller product
		}

		// Admin Customers
		adminCustomers := admin.Group("/customers")
		{
			adminCustomers.GET("/", controllers.GetAllCustomers)      // Pastikan GetAllCustomers ada di controller customer
			adminCustomers.PUT("/:id", controllers.UpdateCustomer)    // Pastikan UpdateCustomer ada di controller customer
			adminCustomers.DELETE("/:id", controllers.DeleteCustomer) // Pastikan DeleteCustomer ada di controller customer
		}

		// Admin Users
		adminUsers := admin.Group("/users")
		{
			adminUsers.GET("/", controllers.GetAllUsers)      // Pastikan GetAllUsers ada di controller user
			adminUsers.PUT("/:id", controllers.UpdateUser)    // Pastikan UpdateUser ada di controller user
			adminUsers.DELETE("/:id", controllers.DeleteUser) // Pastikan DeleteUser ada di controller user
		}

		admin.PUT("/orders/:id/deliver", controllers.DeliverOrder) // Pastikan DeliverOrder ada di controller order
	}

	return r
}
