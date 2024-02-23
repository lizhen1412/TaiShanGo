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

	// 配置热重载
	// 当配置文件发生更改时，会触发该函数内的操作。
	// 参数 e 是 fsnotify.Event 类型，表示配置文件的变化事件。
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 打印配置文件已更改的消息和文件名
		fmt.Println("Config file changed:", e.Name)

		// 尝试重新加载配置文件并解析到全局配置变量 GlobalConfig 中
		if err := viper.Unmarshal(&GlobalConfig); err != nil {
			// 如果重新加载配置文件失败，则记录错误信息
			log.Printf("Error re-loading config: %s", err)
		}
	})
}
