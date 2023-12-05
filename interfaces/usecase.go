package interfaces

import (
	"assignment_4/entities/models"
	"context"
)

type (
	UserUsecase interface {
		RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
		LoginUser(ctx context.Context, email, password string) (*models.User, error)
	}

	TransactionUsecase interface{}
)
