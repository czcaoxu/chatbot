package database

import "chatbot/lib"

type DataBase interface {
	QueryHistoryMessages(userID string) []*lib.Message
	SaveMessage(userID, userMsg, botResponse string)
}
