package ioc

import (
	"chatbot/api"
	"chatbot/api/handler"
)

type App struct {
	ChatHandler    *handler.ChatHandler
	ModelsHandler  *handler.ModelsHandler
	HistoryHandler *handler.HistoryHandler
	ClearHandler   *handler.ClearHandler
}

func NewApp(
	chatHandler *handler.ChatHandler,
	modelsHandler *handler.ModelsHandler,
	historyHandler *handler.HistoryHandler,
	clearHandler *handler.ClearHandler,
) *App {
	return &App{
		ChatHandler:    chatHandler,
		ModelsHandler:  modelsHandler,
		HistoryHandler: historyHandler,
		ClearHandler:   clearHandler,
	}
}

func (a *App) Run() error {
	router := api.InitRouter(
		a.ChatHandler,
		a.ModelsHandler,
		a.HistoryHandler,
		a.ClearHandler,
	)
	return router.Run(":8080")
}
