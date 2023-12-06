package apps

import (
	"assignment_4/configs"
	"assignment_4/interfaces"
	"assignment_4/usecases"
)

type Usecase struct {
	interfaces.UserUsecase
	interfaces.TransactionUsecase
	interfaces.GameUsecase
}

func NewUsecase(repo configs.Repository) *Usecase {
	return &Usecase{
		UserUsecase:        usecases.NewUserUsecase(repo.UserRepository, repo.WalletRepository, repo.ForgotPasswordRepository),
		TransactionUsecase: usecases.NewTransactionUsecase(repo.SourceOfFundRepository, repo.WalletRepository, repo.TransactionRepository, repo.UserRepository),
		GameUsecase:        usecases.NewGameUsecase(repo.GameRepository, repo.UserRepository, repo.TransactionRepository, repo.WalletRepository, repo.BoxRepository),
	}
}
