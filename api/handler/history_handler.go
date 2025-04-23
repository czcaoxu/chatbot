package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HistoryHandler struct {
}

func NewHistoryHandler() *HistoryHandler {
	return &HistoryHandler{}
}

func (c *HistoryHandler) GetHistory(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "会话历史记录"})
}
