package main

import (
	"context"
	"os"

	"github.com/lizhen1412/TaiShanGo/config"
	"github.com/lizhen1412/TaiShanGo/database"
	"github.com/lizhen1412/TaiShanGo/logger" // 确保导入了您的logger包
	"go.uber.org/zap"
)

func main() {
	// 设置环境变量，指定当前环境为开发环境
	os.Setenv("APP_ENV", "development")

	// 从环境变量获取当前环境，如果未设置则默认为 development
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// 初始化配置
	config.Initialize(env)

	// 初始化日志系统
	logger.InitLogger("info") // 可以根据配置文件来设置日志级别

	// 使用zap记录一条信息日志，表明配置已初始化
	logger.Info("配置初始化完成", zap.String("环境", env))

	// 从配置文件中获取Redis连接信息
	redisConfig := config.GlobalConfig.Redis

	// 初始化Redis客户端
	redisClient, err := database.InitRedis(&redisConfig)
	if err != nil {
		logger.Error("连接到Redis失败", zap.Error(err))
		os.Exit(1)
	}

	// 使用zap记录一条信息日志，表明Redis客户端已初始化
	logger.Info("Redis客户端初始化完成")

	// 使用Redis客户端设置一个键值对
	ctx := context.Background()
	err = redisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		logger.Error("设置Redis键值对失败", zap.Error(err))
	}

	val, err := redisClient.Get(ctx, "key").Result()
	if err != nil {
		logger.Error("从Redis获取键值对失败", zap.Error(err))
	} else {
		logger.Info("从Redis获取的键值对", zap.String("key", "value"), zap.String("val", val))
	}

	// 示例结束，记录一条日志
	logger.Info("示例执行完毕")
}
