package apps

import (
	"assignment_4/configs"
	"assignment_4/handlers"
	"assignment_4/interfaces"
)

type (
	Handlers struct {
		UserHandler        interfaces.UserHandler
		TransactionHandler interfaces.TransactionHandler
		GameHandler        interfaces.GameHandler
	}
)

func NewHandlers(repo configs.Repository) (*Handlers, error) {

	usecase := NewUsecase(repo)

	return &Handlers{
		UserHandler:        handlers.NewUserHandler(usecase.UserUsecase),
		TransactionHandler: handlers.NewTransactionHandler(usecase.TransactionUsecase),
		GameHandler:        handlers.NewGameHandler(usecase.GameUsecase),
	}, nil
}
