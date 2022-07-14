package main

import (
	"github.com/gin-gonic/gin"

	"go-gin/routes"
	"go-gin/config"
)

func main() {
	r := InitServer()

	r.Run(":3000")
}

func InitServer() *gin.Engine {
	// init router
	router := gin.Default()

	// init db
	db := config.InitDB()

	APIv1 := router.Group("/api/v1")
	{
		routes.RoutesMahasiswa(db, APIv1)
	}

	return router
}