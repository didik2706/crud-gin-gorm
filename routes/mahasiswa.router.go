package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-gin/controllers"
)

func RoutesMahasiswa(db *gorm.DB, c *gin.RouterGroup) {
	c.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	mahasiswa := c.Group("/mahasiswa")
	{
		// Route GET
		mahasiswa.GET("/", controllers.FindAllMahasiswa)
		mahasiswa.GET("/:id", controllers.FindOneMahasiswa)

		// Route POST
		mahasiswa.POST("/", controllers.CreateMahasiswa)

		// Route PUT
		mahasiswa.PUT("/:id", controllers.UpdateMahasiswa)

		// Route DELETE
		mahasiswa.DELETE("/:id", controllers.DeleteMahasiswa)
	}
}