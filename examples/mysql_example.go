package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lizhen1412/TaiShanGo/config"
	"github.com/lizhen1412/TaiShanGo/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User 结构体表示用户模型
type User struct {
	ID    uint   // 用户ID
	Name  string // 用户名
	Email string // 用户邮箱
}

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

	// 从配置文件中获取数据库连接信息
	dbConfig := config.GlobalConfig.Database

	// 构建数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info), // 使用自定义Logger
		PrepareStmt: true,                                // 启用预编译语句
	})

	if err != nil {
		panic("failed to connect database")
	}

	// 获取底层sql.DB对象以设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database")
	}

	// 设置最大打开的连接数
	sqlDB.SetMaxOpenConns(100)

	// 设置最大闲置的连接数
	sqlDB.SetMaxIdleConns(25)

	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

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
