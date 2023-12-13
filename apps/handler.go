package apps

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/configs"
	rest "git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/handlers/rest"
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
		UserHandler:        rest.NewUserHandler(usecase.UserUsecase),
		TransactionHandler: rest.NewTransactionHandler(usecase.TransactionUsecase),
		GameHandler:        rest.NewGameHandler(usecase.GameUsecase),
	}, nil
}

// func NewGrpcHandler(repo configs.Repository) (*)
