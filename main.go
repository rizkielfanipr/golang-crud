package main

import (
	"log"

	"golang-crud/config" // pastikan path sesuai dengan struktur folder Anda
	"golang-crud/routes" // pastikan path sesuai dengan struktur folder Anda

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Menghubungkan database
	config.ConnectDatabase()

	// Inisialisasi router Gin
	r := gin.Default()

	// Konfigurasi CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Ganti dengan URL frontend Anda
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true, // Jika Anda ingin menggunakan cookies
	}))

	// Setup routing
	routes.SetupRouter()

	// Menjalankan server pada port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
