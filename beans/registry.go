package beans

import (
	"ginDemo/config"
	"ginDemo/controller"
	"ginDemo/dao"
	"ginDemo/service/impl"
)

// 注册所有的beans
func RegistryBeans() {
	config.Register("sheepDao", dao.NewSheepDao)
	config.Register("sheepService", impl.NewSheepServiceImpl)
	config.Register("sheepController", controller.NewSheepController)
}
