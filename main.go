package main

import (
	"latihan-go-crud/config"
	"latihan-go-crud/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.InitDB()

	routes.SiswaRoutes(router)

	// Fungsi yang diekspor sebagai titik masuk aplikasi
	Handler := router.Handler()

	// Memulai server HTTP menggunakan fungsi Handler yang diekspor
	if err := http.ListenAndServe(":8080", Handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
