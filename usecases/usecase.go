package usecases

import (
	"assignment_4/configs"
	"assignment_4/interfaces"
)

type Usecase struct {
	interfaces.UserUsecase
	interfaces.TransactionUsecase
}

func NewUsecase(repo configs.Repository) *Usecase {
	return &Usecase{
		UserUsecase:        NewUserUsecase(repo.UserRepository, repo.WalletRepository),
		TransactionUsecase: NewTransactionUsecase(repo.SourceOfFundRepository, repo.WalletRepository, repo.TransactionRepository),
	}
}
