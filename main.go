package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harish1907/zgolang/controllers"
	"github.com/harish1907/zgolang/docs"
	"github.com/harish1907/zgolang/initializers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LocalEnvironmentVariable()
	initializers.ConnectionDB()
	initializers.MigrateDatabaseSystem()
}

func main() {
	r := gin.Default()
	r.POST("singup/user/", controllers.SignUp)
	r.POST("login/user/", controllers.LoginAPI)

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}