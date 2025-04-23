//go:build wireinject

package ioc

import (
	"chatbot/api/handler"
	"chatbot/internal/ioc"
	"chatbot/internal/repository"
	"chatbot/internal/repository/dao"
	"chatbot/internal/service/dialogue"
	"github.com/google/wire"
)

var (
	BaseSet = wire.NewSet(
		ioc.InitDB,
		ioc.InitLLMClient,
	)

	dialogueSvcSet = wire.NewSet(
		dao.NewDialogueDAO,
		repository.NewDialogueRepository,
		dialogue.NewDialogueService,
	)

	handlersSet = wire.NewSet(
		handler.NewChatHandler,
		handler.NewModelsHandler,
		handler.NewHistoryHandler,
		handler.NewClearHandler,
	)
)

func InitApp() *ioc.App {
	wire.Build(
		BaseSet,

		dialogueSvcSet,

		handlersSet,

		ioc.NewApp,
	)

	return nil
}
