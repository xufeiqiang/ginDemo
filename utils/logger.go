package utils

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
	"time"
)

// 全局的日志实例
var loggerInstance *logrus.Logger
var once sync.Once

// InitLogger 初始化全局日志实例
func InitLogger() {
	once.Do(func() {
		// 创建一个新的 Logrus 实例
		logger := logrus.New()

		// 设置日志级别为 Info
		logger.SetLevel(logrus.InfoLevel)

		// 根据时间戳生成日志文件名
		currentTime := time.Now().Format("2006-01-02_15-04-05.000")
		logFileName := fmt.Sprintf("logs/app_log_%s.log", currentTime)

		// 配置日志轮换功能
		fileWriter := &lumberjack.Logger{
			Filename:   logFileName,
			MaxSize:    10, // MB
			MaxBackups: 3,
			MaxAge:     7, // days
			Compress:   true,
		}
		// 多路输出：文件 + 控制台
		multiWriter := io.MultiWriter(os.Stdout, fileWriter)
		logger.SetOutput(multiWriter)
		// 设置日志格式为 text 格式
		logger.SetFormatter(&logrus.TextFormatter{})
		//logger.SetFormatter(&logrus.JSONFormatter{})

		// 保存为全局变量
		loggerInstance = logger
	})
}

// GetLogger 获取全局日志实例
func GetLogger() *logrus.Logger {
	return loggerInstance
}
