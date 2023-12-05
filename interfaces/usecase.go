package interfaces

import (
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/entities/payload/res"
	"context"
)

type (
	UserUsecase interface {
		RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
		LoginUser(ctx context.Context, login *req.UserLoginRequest) (*res.UserLoginResponse, error)
	}

	TransactionUsecase interface {
		CreateTransferTransaction(ctx context.Context, transfer *req.TransferRequest) (*models.Transaction, error)
		CreateTopupTransaction(ctx context.Context, topup *req.TopupRequest) (*models.Transaction, error)
	}
)
