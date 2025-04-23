package dao

//import "fmt"

//import (
//	"chatbot/internal/domain"
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"os"
//)
//
//// MySQLClient
//type MySQLClient struct {
//	client *gorm.DB
//}
//
//// NewMySQLClient
//func NewMySQLClient() DataBase {
//	client := connectDB()
//	return &MySQLClient{
//		client: client,
//	}
//}
//
//// 连接数据库
//func connectDB() *gorm.DB {
//	dbHost := os.Getenv("MYSQL_HOST")         // 读取 MySQL 主机
//	dbPort := os.Getenv("MYSQL_PORT")         // 读取 MySQL 端口
//	dbUser := os.Getenv("MYSQL_USER")         // 读取 MySQL 用户
//	dbPassword := os.Getenv("MYSQL_PASSWORD") // 读取 MySQL 密码
//	dbName := os.Getenv("MYSQL_DATABASE")     // 读取数据库名
//
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//		dbUser, dbPassword, dbHost, dbPort, dbName)
//
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Println("数据库连接失败:", err)
//		return nil
//	}
//
//	// 自动迁移表结构
//	db.AutoMigrate(&Message{})
//	fmt.Println("数据库连接成功！")
//	return db
//}
//
//// 保存聊天记录
//func (m *MySQLClient) CreateDialogue(req *domain.ChatRequest, botResponse string) {
//	message := Message{
//		UserID:    req.UserID,
//		SessionID: req.SessionID,
//		Model:     req.Model,
//		Message:   req.Text,
//		Response:  botResponse}
//	m.client.Create(&message)
//	fmt.Println("聊天记录已保存！")
//}

//func (m *MySQLClient) QueryDialogue(req *domain.ChatRequest) []*Message {
//	messages := []*Message{}
//	result := m.client.Where("user_id = ? AND session_id = ? AND model = ?", req.UserID, req.SessionID, req.Model).
//		Order("created_at asc").
//		Find(&messages)
//
//	if result.Error != nil {
//		fmt.Println("查寻出错：", result.Error)
//		return nil
//	}
//
//	return messages
//}
