package apps

import (
	"assignment_4/configs"
	"assignment_4/handlers"
	"assignment_4/interfaces"
	"assignment_4/usecases"
)

type (
	Handlers struct {
		UserHandler interfaces.UserHandler
	}
)

func NewHandlers(repo configs.Repository) (*Handlers, error) {

	usecase := usecases.NewUsecase(repo)

	return &Handlers{
		UserHandler: handlers.NewUserHandler(usecase.UserUsecase),
	}, nil
}
