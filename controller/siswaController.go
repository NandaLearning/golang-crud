package controller

import (
	"fmt"
	"latihan-go-crud/config"
	"latihan-go-crud/models"

	"github.com/gin-gonic/gin"
)

func GetSiswa(router *gin.Context) {
	var siswa []models.Siswa

	db := config.GetDB()

	if err := db.Find(&siswa).Error; err != nil {
		fmt.Println("ada yang error")
	}

	router.JSON(200, gin.H{
		"data": siswa,
	})
}

func GetSiswaById(router *gin.Context) {
	id := router.Param("id")

	var siswa models.Siswa
	db := config.GetDB()

	if err := db.First(&siswa, "id = ?", id).Error; err != nil {
		fmt.Println("error di id bang")
	}

	router.JSON(200, gin.H{
		"message": "berhasil",
		"status":  200,
		"data":    siswa,
	})
}

func CreateSiswa(router *gin.Context) {
	var siswa models.Siswa

	if err := router.ShouldBindBodyWithJSON(&siswa); err != nil {
		fmt.Println("error di pertama create")
	}

	db := config.GetDB()

	if err := db.Create(&siswa).Error; err != nil {
		fmt.Println("error di kedua create")
	}

	router.JSON(200, gin.H{
		"message": "berhasil membuat data",
		"status":  200,
		"data":    siswa,
	})
}

func UpdateSiswa(router *gin.Context) {
	id := router.Param("id")

	var siswa models.Siswa
	db := config.GetDB()

	if err := db.First(&siswa, "id = ?", id).Error; err != nil {
		fmt.Println("error bang")
	}

	if err := router.ShouldBindBodyWithJSON(&siswa); err != nil {
		fmt.Println("binding gagal")
	}

	if err := db.Save(&siswa).Error; err != nil {
		fmt.Println("save data gagal")
	}

	router.JSON(200, gin.H{
		"message": "update data berhasil",
		"status":  200,
		"data":    siswa,
	})
}

func DeleteSiswaById(router *gin.Context) {
	var siswa models.Siswa
	id := router.Param("id")
	db := config.GetDB()
	if err := db.First(&siswa, "id = ?", id).Error; err != nil {
		fmt.Println("gagal mendapatkan nama siswa")
	}

	if err := db.Delete(&siswa).Error; err != nil {
		fmt.Println("gagal menghapus akun")
	}

	router.JSON(200, gin.H{
		"message": "akun berhasil di hapus",
		"status":  200,
		"data":    siswa,
	})
}

func DeleteSiswaAll(router *gin.Context) {
	db := config.GetDB()

	if err := db.Exec("DELETE FROM siswas").Error; err != nil {
		fmt.Println(err)
	}

	router.JSON(200, gin.H{
		"message": "semua data berhasil di hapus",
		"status":  200,
	})
}
