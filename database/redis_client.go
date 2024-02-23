package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/lizhen1412/TaiShanGo/config"
)

// InitRedis 根据配置初始化并返回一个Redis客户端实例。
// 这样设计允许调用者更灵活地处理错误和管理Redis连接。
func InitRedis(cfg *config.RedisConfig) (*redis.Client, error) {
	// 根据配置创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), // Redis地址和端口
		Password: cfg.Password,                             // Redis密码，如果没有设置密码，则可以为空
		DB:       cfg.DB,                                   // 使用的Redis数据库索引，默认为0
	})

	// 使用背景上下文测试Redis连接
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("无法连接到Redis: %w", err)
	}

	fmt.Println("成功连接到Redis")
	return client, nil
}
