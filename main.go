package main

import (
	"chatbot/ai"
	"context"
	"database/sql"
	"github.com/go-redis/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	router = ai.NewModelRouter()
	db     *sql.DB
	rdb    *redis.Client
)

// 聊天 API
func chatHandler(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id"`
		Text   string `json:"text"`
		Model  string `json:"model"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	model, err := router.GetModel(req.Model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的模型"})
		return
	}

	// 读取历史对话（上下文）
	//historyKey := "chat:" + req.UserID
	//prevMessages, _ := rdb.Get(historyKey).Result()
	prevMessages := ""

	reply, err := model.Chat(context.Background(), prevMessages+req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI 处理失败, " + err.Error()})
		return
	}

	// 记录聊天历史
	//rdb.Set(historyKey, prevMessages+req.Text+"\nAI: "+reply, 10*time.Minute)
	//db.Exec("INSERT INTO chat_history (user_id, message, response) VALUES (?, ?, ?)", req.UserID, req.Text, reply)

	c.JSON(http.StatusOK, gin.H{"reply": reply})
}

// 获取支持的模型
func modelsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"models": []string{"openai", "qwen"}})
}

// 获取聊天历史
func historyHandler(c *gin.Context) {
	userID := c.Query("user_id")
	rows, _ := db.Query("SELECT message, response FROM chat_history WHERE user_id = ? ORDER BY created_at DESC LIMIT 20", userID)

	var history []map[string]string
	for rows.Next() {
		var msg, resp string
		rows.Scan(&msg, &resp)
		history = append(history, map[string]string{"user": msg, "bot": resp})
	}
	c.JSON(http.StatusOK, history)
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
	r.POST("/chat", chatHandler)
	r.GET("/models", modelsHandler)
	r.GET("/history", historyHandler)
	r.POST("/clear", clearHandler)

	r.Run(":8080")
}
