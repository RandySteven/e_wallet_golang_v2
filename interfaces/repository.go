package interfaces

import (
	"assignment_4/entities"
	"assignment_4/entities/models"
	"context"
)

type (
	Repository[T any, I any] interface {
		Save(ctx context.Context, entity *T) (*T, error)
		FindAll(ctx context.Context) ([]T, error)
		GetById(ctx context.Context, id uint) (*T, error)
		Update(ctx context.Context, entity *T) (*T, error)
		BeginTrx(ctx context.Context) I
		CommitOrRollback(ctx context.Context)
		Count(ctx context.Context) (uint, error)
	}
	UserRepository interface {
		Repository[models.User, UserRepository]
		GetByEmail(ctx context.Context, email string) (*models.User, error)
		RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
	}

	ForgotPasswordRepository interface {
		Repository[models.ForgotPasswordToken, ForgotPasswordRepository]
		GetPasswordTokenByToken(ctx context.Context, token string) (*models.ForgotPasswordToken, error)
		UpdateUserPassword(ctx context.Context, token *models.ForgotPasswordToken, password string) (*models.User, error)
	}

	WalletRepository interface {
		Repository[models.Wallet, WalletRepository]
		GetByNumber(ctx context.Context, number string) (*models.Wallet, error)
		GetByUserId(ctx context.Context, id uint) (*models.Wallet, error)
	}

	TransactionRepository interface {
		Repository[models.Transaction, TransactionRepository]
		CreateTopupTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
		CreateTransferTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
		GetTransactionsByWalletId(ctx context.Context, walletId uint) ([]models.Transaction, error)
		GetAllTransactions(ctx context.Context, query *entities.QueryCondition) ([]models.Transaction, error)
	}

	SourceOfFundRepository interface {
		Repository[models.SourceOfFund, SourceOfFundRepository]
		GetSourceOfFundBySource(ctx context.Context, source string) (*models.SourceOfFund, error)
	}

	GameRepository interface {
		Repository[models.Game, GameRepository]
	}

	BoxRepository interface {
		Repository[models.Box, BoxRepository]
		GetNineRandomBoxes(ctx context.Context) ([]models.Box, error)
	}
)
