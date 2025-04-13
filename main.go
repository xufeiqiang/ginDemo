package main

import (
	"ginDemo/beans" // 引入 beans 包
	"ginDemo/config"
	"ginDemo/router"
	"ginDemo/utils"
)

func main() {
	utils.InitLogger()
	logger := utils.GetLogger()

	logger.Info("---------1. 注册 Bean ------------")
	beans.RegistryBeans() //  注册所有 Bean（在容器初始化前）

	logger.Info("---------2. 初始化 Bean 容器 ------------")
	config.InitApplicationContext()

	logger.Info("---------3. 初始化数据库 ------------")
	config.InitDB()

	logger.Info("---------4. 启动服务 ------------")
	r := router.SetupRouter()
	r.Run(":8080")
}
