package usecases

import (
	"assignment_4/entities/models"
	"assignment_4/interfaces"
	"context"
)

type userUsecase struct {
	userRepo   interfaces.UserRepository
	walletRepo interfaces.WalletRepository
}

// LoginUser implements interfaces.UserUsecase.
func (usecase *userUsecase) LoginUser(ctx context.Context, email string, password string) (*models.User, error) {
	panic("unimplemented")
}

// RegisterUser implements interfaces.UserUsecase.
func (usecase *userUsecase) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := usecase.userRepo.RegisterUser(ctx, user)
	return user, err
}

func NewUserUsecase(userRepo interfaces.UserRepository, walletRepo interfaces.WalletRepository) *userUsecase {
	return &userUsecase{
		userRepo:   userRepo,
		walletRepo: walletRepo,
	}
}

var _ interfaces.UserUsecase = &userUsecase{}
