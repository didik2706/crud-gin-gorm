package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"go-gin/models"
)

func InitDB() *gorm.DB {
	dsn := "didik27:Didik.,.88@tcp(127.0.0.1:3306)/mahasiswa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.Migrator().CreateTable(&models.Mahasiswa{})

	return db
}