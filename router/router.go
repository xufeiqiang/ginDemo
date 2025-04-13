package router

import (
	"ginDemo/config"
	"ginDemo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	engine := gin.Default()
	cxf := config.GetApplicationContext()
	sheepController := cxf.GetBean("sheepController").(*controller.SheepController)
	userGroup := engine.Group("/sheep")
	{
		userGroup.GET("/:value", sheepController.GetSheepByValue)
		userGroup.POST("/insert", sheepController.InsertSheep)
		userGroup.POST("/list", sheepController.GetSheepList)
		//userGroup.POST("/", userController.CreateUser)
	}
	return engine
}
