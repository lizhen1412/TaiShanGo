package main

import (
	"fmt"
	"os"

	"github.com/lizhen1412/TaiShanGo/config" // 替换为你的模块路径
)

func main() {
	// 设置环境变量，这通常是在你的应用启动脚本中设置，或者根据部署环境预设
	// 为了示例，我们在这里手动设置
	os.Setenv("APP_ENV", "development")

	// 从环境变量获取当前环境，如果未设置则默认为 development
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// 初始化配置
	config.Initialize(env)

	// 访问配置信息
	dbConfig := config.GlobalConfig.Database
	fmt.Println("Database Configuration:")
	fmt.Printf("Host: %s\nPort: %d\nUser: %s\nPassword: %s\nName: %s\n",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

	redisConfig := config.GlobalConfig.Redis
	fmt.Println("Redis Configuration:")
	fmt.Printf("Host: %s\nPort: %d\nPassword: %s\nDB: %d\n",
		redisConfig.Host, redisConfig.Port, redisConfig.Password, redisConfig.DB)

	// 验证配置有效性
	if err := config.GlobalConfig.Validate(); err != nil {
		fmt.Println("Configuration validation error:", err)
		os.Exit(1)
	}

	// 使用配置信息初始化数据库连接、Redis 客户端等...
	// 此处略过实际的初始化逻辑，仅展示如何使用配置信息
}
