package database

import "chatbot/lib"

type DataBase interface {
	QueryHistoryMessages(req *lib.ChatRequest) []*lib.Message
	SaveMessage(req *lib.ChatRequest, botResponse string)
}
