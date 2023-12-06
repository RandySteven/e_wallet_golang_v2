package interfaces

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/req"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/res"
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
		GetAllTransactionsRecords(ctx context.Context, query *entities.QueryCondition, userId uint) (*res.TransactionPaginationResponses, error)
	}

	GameUsecase interface {
		PlayGame(ctx context.Context, game *models.Game) (*models.Game, error)
		ChooseReward(ctx context.Context, chooseReward *req.ChooseReward) (*models.Game, error)
		GetUserCurrentChance(ctx context.Context, userId uint) (*models.User, error)
	}
)
