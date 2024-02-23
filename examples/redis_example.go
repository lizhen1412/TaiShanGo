package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lizhen1412/TaiShanGo/config"
	"github.com/lizhen1412/TaiShanGo/database"
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

	// 从配置文件中获取Redis连接信息
	redisConfig := config.GlobalConfig.Redis

	// 初始化Redis客户端
	redisClient, err := database.InitRedis(&redisConfig)
	if err != nil {
		panic("failed to connect to Redis: " + err.Error())
	}

	// 使用Redis客户端设置一个键值对
	ctx := context.Background()
	err = redisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic("failed to set key in Redis: " + err.Error())
	}

	// 获取并打印出之前设置的键的值
	val, err := redisClient.Get(ctx, "key").Result()
	if err != nil {
		panic("failed to get key from Redis: " + err.Error())
	}
	fmt.Printf("The value of 'key' is: %s\n", val)

	// redis 总数
	dbSize, err2 := redisClient.DBSize(ctx).Result()
	if err2 != nil {
		panic("failed to get db size from Redis: " + err2.Error())
	}
	fmt.Printf("The db size is: %d\n", dbSize)

	// redis 随机key
	randomKey, err3 := redisClient.RandomKey(ctx).Result()
	if err3 != nil {
		panic("failed to get random key from Redis: " + err3.Error())
	}

	fmt.Printf("The random key is: %s\n", randomKey)

	// 删除 set 的 key
	err4 := redisClient.Del(ctx, "key").Err()
	if err4 != nil {
		panic("failed to delete key from Redis: " + err4.Error())
	}

}
