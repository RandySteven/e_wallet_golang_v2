package interfaces

import (
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
	}
	UserRepository interface {
		Repository[models.User, UserRepository]
		GetByEmail(ctx context.Context, email string) (*models.User, error)
		RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
	}

	WalletRepository interface {
		Repository[models.Wallet, WalletRepository]
	}

	TransactionRepository interface {
		Repository[models.Transaction, TransactionRepository]
	}
)
