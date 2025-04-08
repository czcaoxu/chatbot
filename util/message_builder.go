package util

import (
	"chatbot/lib"
)

func BuildHistoricalMessages(input []*lib.Message) []map[string]string {
	var context []map[string]string
	for _, msg := range input {
		// 用户的消息
		context = append(context, map[string]string{
			"role":    "user",
			"content": msg.Message,
		})

		// AI 的回复
		context = append(context, map[string]string{
			"role":    "assistant",
			"content": msg.Response,
		})
	}
	return context
}
