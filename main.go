package main

import (
	"chatbot/ai"
	"chatbot/database"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

var (
	//cfg, _ = config.LoadConfig("./config/service.json")
	router = ai.NewModelRouter()
	db     = database.NewMySQLClient()
	rdb    *redis.Client
)

// 聊天 API

// 获取支持的模型
func modelsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"models": []string{"openai", "qwen"}})
}

// 获取聊天历史
func historyHandler(c *gin.Context) {
	//userID := c.Query("user_id")
	//rows, _ := db.Query("SELECT message, response FROM chat_history WHERE user_id = ? ORDER BY created_at DESC LIMIT 20", userID)
	//
	//var history []map[string]string
	//for rows.Next() {
	//	var msg, resp string
	//	rows.Scan(&msg, &resp)
	//	history = append(history, map[string]string{"user": msg, "bot": resp})
	//}
	//c.JSON(http.StatusOK, history)
}

// 清除上下文
func clearHandler(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id"`
	}
	c.ShouldBindJSON(&req)

	rdb.Del("chat:" + req.UserID)
	c.JSON(http.StatusOK, gin.H{"message": "会话已清除"})
}

func main() {
	r := gin.Default()
	r.POST("/chat", ChatHandler)
	r.GET("/models", modelsHandler)
	r.GET("/history", historyHandler)
	r.POST("/clear", clearHandler)

	r.Run(":8080")
}
