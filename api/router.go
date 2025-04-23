package api

import (
	"chatbot/api/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(
	chatHandler *handler.ChatHandler,
	modelsHandler *handler.ModelsHandler,
	historyHandler *handler.HistoryHandler,
	clearHandler *handler.ClearHandler,
) *gin.Engine {

	r := gin.Default()

	r.POST("/chat", chatHandler.HandleChat)
	r.GET("/models", modelsHandler.GetModels)
	r.GET("/history", historyHandler.GetHistory)
	r.POST("/clear", clearHandler.HandleClear)

	return r
}
