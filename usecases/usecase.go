package usecases

import (
	"assignment_4/configs"
	"assignment_4/interfaces"
)

type Usecase struct {
	interfaces.UserUsecase
	interfaces.TransactionUsecase
	interfaces.GameUsecase
}

func NewUsecase(repo configs.Repository) *Usecase {
	return &Usecase{
		UserUsecase:        NewUserUsecase(repo.UserRepository, repo.WalletRepository, repo.ForgotPasswordRepository),
		TransactionUsecase: NewTransactionUsecase(repo.SourceOfFundRepository, repo.WalletRepository, repo.TransactionRepository, repo.UserRepository),
	}
}
