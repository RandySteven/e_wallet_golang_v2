package usecases

import (
	"assignment_4/configs"
	"assignment_4/interfaces"
)

type Usecase struct {
	interfaces.UserUsecase
}

func NewUsecase(repo configs.Repository) *Usecase {
	return &Usecase{
		UserUsecase: NewUserUsecase(repo.UserRepository, repo.WalletRepository),
	}
}
