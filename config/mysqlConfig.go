package config

import (
	"ginDemo/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

type MysqlConfig struct{}

var db *gorm.DB //
var dbOnce sync.Once

func InitDB() {
	dbOnce.Do(func() {
		utils.GetLogger().Info("开始初始化数据库连接...")
		dsn := "root:root@tcp(localhost:3306)/sheep?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			utils.GetLogger().Fatalf("连接数据库失败,请查看数据库账号密码是否错误: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			utils.GetLogger().Fatalf("获取数据库连接失败: %v", err)
		}

		// 配置连接池参数
		sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
		sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
		sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生命周期
		utils.GetLogger().Info("初始化数据库连接成功...")
	})
}

// 获取 数据库连接实例
func GetMysqlDb() *gorm.DB {
	if db == nil {
		utils.GetLogger().Fatalf("请先执行InitDB方法初始化数据库连接...")
	}
	return db
}
