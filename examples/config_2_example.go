package main

import (
	"fmt"
	"os"

	"github.com/lizhen1412/TaiShanGo/config"
	"github.com/spf13/viper"
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

	// 直接从配置文件中获取 user 和 order 的地址
	userAPI := viper.GetString("api.user")
	orderAPI := viper.GetString("api.order")

	// 打印 user 和 order 的地址
	fmt.Println("User API:", userAPI)
	fmt.Println("Order API:", orderAPI)

}
