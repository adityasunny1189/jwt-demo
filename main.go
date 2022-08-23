package main

import (
	"fmt"

	"github.com/adityasunny1189/jwt-demo/controllers"
	"github.com/adityasunny1189/jwt-demo/database"
	"github.com/adityasunny1189/jwt-demo/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	database.Connect("root:Adisunny123@tcp(127.0.0.1:3306)/jwt_demo_db?charset=utf8mb4&parseTime=True&loc=Local")
	database.Migrate()
}

func main() {
	fmt.Println("Go JWT auth")

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
