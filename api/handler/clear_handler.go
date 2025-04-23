package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClearHandler struct {
}

func NewClearHandler() *ClearHandler {
	return &ClearHandler{}
}

func (c *ClearHandler) HandleClear(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "会话已清除"})
}
