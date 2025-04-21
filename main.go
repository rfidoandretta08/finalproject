package main

import (
	"finalproject/config"
	"finalproject/routes"
)

func main() {
	// Inisialisasi konfigurasi dan database
	config.InitDB()

	// Inisialisasi router
	router := routes.SetupRouter()

	// Jalankan server di port 8080
	router.Run(":8080")
}
