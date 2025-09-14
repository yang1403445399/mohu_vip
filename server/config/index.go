package config

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func FiberRun(app *fiber.App, port string) {
	// 创建一个信号通道，用于监听系统中断信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// 在 goroutine 中启动服务器
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	log.Printf("Server is running on :%s\n", port)
	log.Printf("Use CTRL+C to shutdown\n")

	// 等待中断信号
	<-c

	// 创建超时上下文用于优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	}

	log.Println("Server shutdown successfully")
}

func InitDB() {
	// 自定义日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 日志输出目标
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:                  logger.Info,            // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略记录未找到错误
			Colorful:                  true,                   // 启用彩色日志
		},
	)

	// 打开 mysql 数据库连接，添加连接超时参数
	// timeout: 连接超时时间
	// readTimeout: 读取数据超时时间
	dsn := "root:123456@tcp(127.0.0.1:3306)/mohu_vip?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s&readTimeout=30s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tp_",
			SingularTable: true,
		},
		Logger: newLogger, // 设置日志记录器
	})
	if err != nil {
		log.Fatal(err)
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// 数据库连接池配置优化
	// 设置最大打开连接数 - 可根据应用负载调整
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数 - 建议为最大连接数的10%-20%
	sqlDB.SetMaxIdleConns(20)
	// 设置连接最大存活时间 - 避免连接过长时间不活跃
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	// 设置连接最大空闲时间 - 回收长时间不使用的空闲连接
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	DB = db
}
