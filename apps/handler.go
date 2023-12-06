package apps

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/handlers"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
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
