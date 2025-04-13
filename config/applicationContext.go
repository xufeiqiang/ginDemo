package config

import (
	"ginDemo/utils"
	"sort"
	"sync"
)

type ApplicationContext struct {
	/*
		key bean的name
		val bean的实例化对象
	*/
	beans map[string]interface{}
}

// 定义构造函数类型
type Constructor func(ctx *ApplicationContext) interface{}

/*
注册表
key bean的name
val Constructor 类型
*/
type beanDefinition struct {
	order       int
	name        string
	constructor Constructor
}

var registry = []beanDefinition{} // 改为切片，支持排序

var once sync.Once

// 定义全局变量 applicationContext
var applicationContext *ApplicationContext

// 注册方法
var registerSeq = 0 // 全局自增顺序，用来控制注册顺序

func Register(name string, constructor Constructor) {
	registerSeq++
	registry = append(registry, beanDefinition{
		order:       registerSeq,
		name:        name,
		constructor: constructor,
	})
}

// 初始化所有的bean,按照顺序初始化，例如 impl 依赖于dao,如果先初始化dao会失败
func InitApplicationContext() *ApplicationContext {
	once.Do(func() {
		utils.GetLogger().Info("开始实例化所有的bean")

		applicationContext = &ApplicationContext{
			beans: make(map[string]interface{}),
		}

		// 按照注册顺序排序
		sort.Slice(registry, func(i, j int) bool {
			return registry[i].order < registry[j].order
		})

		// 遍历排序后的列表，初始化所有的bean
		for _, def := range registry {
			bean := def.constructor(applicationContext)
			applicationContext.beans[def.name] = bean
		}
		ShowAllRegisteredBeans()
		utils.GetLogger().Info("实例化所有的bean成功！！！")
	})

	return applicationContext
}

// 根据bean的名称获取bean的实例
func (cxf *ApplicationContext) GetBean(name string) interface{} {
	return cxf.beans[name]
}

// 获取 applicationContext 实例
func GetApplicationContext() *ApplicationContext {
	if applicationContext == nil {
		utils.GetLogger().Fatalf("请先执行InitApplicationContext方法初始化所有的bean...")
	}
	return applicationContext
}

// 显示所有注册的Bean
func ShowAllRegisteredBeans() {
	utils.GetLogger().Info("已注册的Bean列表:")
	for _, def := range registry {
		utils.GetLogger().Infof("Bean name: %s, 注册顺序: %d", def.name, def.order)
	}
}
