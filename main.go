package main

import (
	"finalproject/config" // <- ini yang sebelumnya belum diimport
	"finalproject/routes"
)

func main() {
	// Inisialisasi konfigurasi dan database
	config.InitDB() // Pastikan nama fungsinya sesuai dengan yang kamu buat di config/database.go

	// Inisialisasi router
	router := routes.SetupRouter()

	// Jalankan server di port 8080
	router.Run(":8080")
}
