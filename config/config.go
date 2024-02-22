package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GlobalConfig AppConfig // 全局配置变量

// Initialize 初始化配置系统
func Initialize(env string) {
	viper.AddConfigPath("config") // 配置文件路径
	viper.SetConfigName(env)      // 根据环境加载不同配置文件
	viper.AutomaticEnv()          // 环境变量覆盖

	// 尝试读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 映射配置到结构体
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}

	// 验证配置
	if err := GlobalConfig.Validate(); err != nil {
		log.Fatalf("Configuration validation error: %s", err)
	}

	// 配置热重载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(&GlobalConfig); err != nil {
			log.Printf("Error re-loading config: %s", err)
		}
	})
}
