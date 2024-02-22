package main

import (
	"fmt"
	"os"

	"github.com/lizhen1412/TaiShanGo/config"
	"github.com/lizhen1412/TaiShanGo/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func main() {

	os.Setenv("APP_ENV", "development")

	// 从环境变量获取当前环境，如果未设置则默认为 development
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// 初始化配置
	config.Initialize(env)

	// 从配置文件中获取数据库连接信息
	dbConfig := config.GlobalConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// 使用自定义连接池设置
		PrepareStmt: true, // 启用预编译语句
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 执行数据库迁移
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database")
	}

	userRepo := &database.GormRepository[User]{}

	// 示例：创建一个用户
	newUser := User{Name: "John Doe", Email: "john@example.com"}
	err = userRepo.Create(db, &newUser)
	if err != nil {
		panic("failed to create user")
	}
	fmt.Println("New user ID:", newUser.ID)
}
