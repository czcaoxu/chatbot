package ioc

import (
	"chatbot/internal/repository/dao"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitDB() *gorm.DB {
	return connectDB()
}

// 连接数据库
func connectDB() *gorm.DB {
	dbHost := os.Getenv("MYSQL_HOST")         // 读取 MySQL 主机
	dbPort := os.Getenv("MYSQL_PORT")         // 读取 MySQL 端口
	dbUser := os.Getenv("MYSQL_USER")         // 读取 MySQL 用户
	dbPassword := os.Getenv("MYSQL_PASSWORD") // 读取 MySQL 密码
	dbName := os.Getenv("MYSQL_DATABASE")     // 读取数据库名

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		panic(err)
	}

	// 自动迁移表结构
	if err := dao.InitTables(db); err != nil {
		fmt.Println("数据库自动迁移表结构失败:", err)
		panic(err)
	}

	fmt.Println("数据库连接成功！")
	return db
}
