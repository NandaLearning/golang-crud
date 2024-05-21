package main

import (
	"latihan-go-crud/config"
	"latihan-go-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.InitDB()

	routes.SiswaRoutes(router)
	router.Run()
}
