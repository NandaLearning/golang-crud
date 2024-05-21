package routes

import (
	"latihan-go-crud/controller"

	"github.com/gin-gonic/gin"
)

func SiswaRoutes(router *gin.Engine) {
	getSiswa := controller.GetSiswa
	getSiswaById := controller.GetSiswaById
	createSiswa := controller.CreateSiswa
	updateSiswa := controller.UpdateSiswa
	deleteSiswaById := controller.DeleteSiswaById
	DeleteSiswaAll := controller.DeleteSiswaAll

	router.GET("/siswa", getSiswa)
	router.GET("/siswa/:id", getSiswaById)
	router.POST("/siswa", createSiswa)
	router.PATCH("/siswa/:id", updateSiswa)
	router.DELETE("/siswa/:id", deleteSiswaById)
	router.DELETE("/siswa", DeleteSiswaAll)
}
