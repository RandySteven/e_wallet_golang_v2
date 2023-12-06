package interfaces

import (
	"assignment_4/entities"
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/entities/payload/res"
	"context"
)

type (
	UserUsecase interface {
		RegisterUser(ctx context.Context, user *models.User) (*models.User, error)
		LoginUser(ctx context.Context, login *req.UserLoginRequest) (*res.UserLoginResponse, error)
		GetUserDetail(ctx context.Context, id uint) (*res.UserDetail, error)
		ForgotPassword(ctx context.Context, forgot *req.ForgotPasswordRequest) (*models.ForgotPasswordToken, error)
		ResetPassword(ctx context.Context, reset *req.PasswordResetRequest) (*models.User, error)
	}

	TransactionUsecase interface {
		CreateTransferTransaction(ctx context.Context, transfer *req.TransferRequest) (*models.Transaction, error)
		CreateTopupTransaction(ctx context.Context, topup *req.TopupRequest) (*models.Transaction, error)
		GetUserHistoryTransactions(ctx context.Context, id uint) ([]models.Transaction, error)
		GetAllTransactionsRecords(ctx context.Context, query *entities.QueryCondition, userId uint) (*res.TransactionPaginationResponses, error)
	}

	GameUsecase interface {
		PlayGame(ctx context.Context, game *models.Game) (*models.Game, error)
		ChooseReward(ctx context.Context, chooseReward *req.ChooseReward) (*models.Game, error)
	}
)
