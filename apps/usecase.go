package apps

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/usecases"
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
