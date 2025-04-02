package database

import (
	"chatbot/lib"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLClient
type MySQLClient struct {
	client *gorm.DB
}

// NewMySQLClient
func NewMySQLClient() DataBase {
	client := connectDB()
	return &MySQLClient{
		client: client,
	}
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
		log.Fatal("数据库连接失败:", err)
	}

	// 自动迁移表结构
	db.AutoMigrate(&lib.Message{})
	fmt.Println("数据库连接成功！")
	return db
}

// 保存聊天记录
func (m *MySQLClient) SaveMessage(userID, userMsg, botResponse string) {
	message := lib.Message{UserID: userID, Message: userMsg, Response: botResponse}
	m.client.Create(&message)
	fmt.Println("聊天记录已保存！")
}

func (m *MySQLClient) QueryHistoryMessages(userID string) []*lib.Message {
	return nil
}
