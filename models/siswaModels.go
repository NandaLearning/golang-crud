package models

import "gorm.io/gorm"

type Siswa struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Nama string `json:"nama"`
	Usia uint   `json:"usia"`
}
