package main

import (
	"echarts_example/internal/controller"
	"echarts_example/internal/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repository.InitPostgresDB()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	r.Use(cors.Default())
	api := r.Group("api")
	{
		api.GET("/chat_one_data", controller.Table{}.GetOneTable)
		api.GET("/chat_all_data", controller.Table{}.GetAllTable)
		api.PUT("/chat_create_data", controller.Table{}.CreateTable)
	}
	r.Run(":4001")

}
