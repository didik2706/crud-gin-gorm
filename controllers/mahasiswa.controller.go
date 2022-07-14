package controllers

import (

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"go-gin/models"
)

type InputCreateMahasiswa struct {
	Name string `json:"name" binding:"required"`
	NIM int `json:"nim" binding:"required"`
	Prodi string `json:"prodi" binding:"required"`
}

var validate *validator.Validate

func FindAllMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mahasiswa []models.Mahasiswa
	db.Find(&mahasiswa)

	c.JSON(200, gin.H{
		"data": mahasiswa,
	})
}

func FindOneMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	
	id := c.Param("id")
	var mahasiswa models.Mahasiswa

	result := db.First(&mahasiswa, "id = ?", id)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "data not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": mahasiswa,
		"affectedRow": result.RowsAffected,
	})
}

func CreateMahasiswa(c *gin.Context) {
	var input InputCreateMahasiswa
	db := c.MustGet("db").(*gorm.DB)

	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	mahasiswa := models.Mahasiswa{Name: input.Name, NIM: input.NIM, Prodi: input.Prodi}
	result := db.Create(&mahasiswa)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(201, gin.H{
		"success": true,
		"message": "data mahasiswa successfully added",
	})
}

func UpdateMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// data param id
	id := c.Param("id")
	var input InputCreateMahasiswa
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	var mahasiswa models.Mahasiswa

	result := db.Find(&mahasiswa, "id = ?", id)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"success": false,
			"message": "data not found",
		})
		return
	}

	// update data
	mahasiswa.Name = input.Name
	mahasiswa.NIM = input.NIM
	mahasiswa.Prodi = input.Prodi
	db.Save(&mahasiswa)

	c.JSON(200, gin.H{
		"success": true,
		"message": "data mahasiswa successfully updated",
	})
}

func DeleteMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// get id from param
	id := c.Param("id")
	// define model
	var mahasiswa models.Mahasiswa
	db.Where("id = ?", id).Delete(&mahasiswa)

	c.JSON(200, gin.H{
		"success": true,
		"message": "data mahasiswa successfully deleted",
	})
}