package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ModelsHandler struct {
}

func NewModelsHandler() *ModelsHandler {
	return &ModelsHandler{}
}

func (c *ModelsHandler) GetModels(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"models": []string{"openai", "qwen"}})
}
