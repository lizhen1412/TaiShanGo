package config

import "fmt"

// DatabaseConfig 定义数据库配置结构体
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`     // 数据库地址
	Port     int    `mapstructure:"port"`     // 数据库端口
	User     string `mapstructure:"user"`     // 数据库用户
	Password string `mapstructure:"password"` // 数据库密码
	Name     string `mapstructure:"name"`     //  数据库名字
}

// RedisConfig 定义Redis配置结构体
type RedisConfig struct {
	Host     string `mapstructure:"host"`     // Redis地址
	Port     int    `mapstructure:"port"`     // Redis端口
	Password string `mapstructure:"password"` // Redis密码
	DB       int    `mapstructure:"db"`       // Redis数据库索引
}

// AppConfig 定义应用配置结构体
type AppConfig struct {
	Database DatabaseConfig `mapstructure:"database"` // 数据库配置
	Redis    RedisConfig    `mapstructure:"redis"`    // Redis配置
}

// Validate 方法用于验证配置的有效性
func (c *AppConfig) Validate() error {
	// 示例验证逻辑，实际应用中需要根据需求定制
	if c.Database.Host == "" || c.Database.Port == 0 {
		return fmt.Errorf("database configuration is invalid") // 数据库配置无效
	}
	if c.Redis.Host == "" || c.Redis.Port == 0 {
		return fmt.Errorf("redis configuration is invalid") // Redis配置无效
	}
	return nil
}
