package main

import (
	"chatbot/lib"
	"chatbot/util"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func ChatHandler(c *gin.Context) {
	req := new(lib.ChatRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	model, err := router.GetModel(req.Model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的模型"})
		return
	}

	if req.SessionID == "" {
		req.SessionID = uuid.New().String()
	}

	// 读取历史对话（上下文）
	historyMessages := db.QueryHistoryMessages(req)
	reply, err := model.Chat(context.Background(), util.BuildHistoricalMessages(historyMessages), req.Text)
	rsp := gin.H{
		"user_id":    req.UserID,
		"session_id": req.SessionID,
	}

	if err != nil {
		rsp["error"] = "AI 处理失败, " + err.Error()
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}

	// 记录聊天历史
	db.SaveMessage(req, reply)

	rsp["reply"] = reply
	c.JSON(http.StatusOK, rsp)
}
