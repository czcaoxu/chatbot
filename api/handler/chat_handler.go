package handler

import (
	"chatbot/internal/domain"
	"chatbot/internal/service/dialogue"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ChatHandler struct {
	dialogueService dialogue.Service
}

func NewChatHandler(svc dialogue.Service) *ChatHandler {
	return &ChatHandler{dialogueService: svc}
}

func (c *ChatHandler) HandleChat(ctx *gin.Context) {
	req := new(ChatRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := c.validateRequest(req); err != nil {
		fmt.Println("request校验失败: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数校验失败"})
		return
	}

	chatInput, err := req.ToDomainEntity()
	if err != nil {
		fmt.Println("request转换失败: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "内部错误"})
		return
	}

	if chatInput.SessionID == "" {
		chatInput.SessionID = uuid.New().String()
	}

	reply, err := c.dialogueService.Chat(ctx, chatInput)
	rsp := gin.H{
		"user_id":    chatInput.UserID,
		"session_id": chatInput.SessionID,
	}

	if err != nil {
		rsp["error"] = "AI 处理失败, " + err.Error()
		ctx.JSON(http.StatusInternalServerError, rsp)
		return
	}

	rsp["reply"] = reply
	ctx.JSON(http.StatusOK, rsp)
}

func (c *ChatHandler) validateRequest(req *ChatRequest) error {
	if !domain.AIModelType(req.ModelType).IsValid() {
		return fmt.Errorf("unsupported model type %v", req.ModelType)
	}

	return nil
}
