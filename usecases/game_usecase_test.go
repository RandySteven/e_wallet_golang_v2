package usecases_test

import (
	"assignment_4/apperror"
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/enums"
	"assignment_4/mocks"
	"assignment_4/usecases"
	"context"
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var boxes = []models.Box{
	{
		ID:     1,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     2,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     3,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     4,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     5,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     6,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     7,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     8,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     9,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     10,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     11,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     12,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     13,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     14,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     15,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     16,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     17,
		Amount: decimal.NewFromInt(50000),
	},
	{
		ID:     18,
		Amount: decimal.NewFromInt(50000),
	},
}

var games = []models.Game{
	{
		ID:       1,
		UserID:   1,
		BoxID1:   2,
		BoxID2:   3,
		BoxID3:   1,
		BoxID4:   12,
		BoxID5:   9,
		BoxID6:   4,
		BoxID7:   5,
		BoxID8:   8,
		BoxID9:   11,
		WinBoxID: 0,
	},
	{
		ID:       2,
		UserID:   1,
		BoxID1:   2,
		BoxID2:   3,
		BoxID3:   1,
		BoxID4:   12,
		BoxID5:   9,
		BoxID6:   4,
		BoxID7:   5,
		BoxID8:   8,
		BoxID9:   11,
		WinBoxID: 1,
	},
}

func TestPlayGame(t *testing.T) {
	t.Run("should return success create game", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}
		gameReq := &models.Game{
			UserID: 1,
		}
		user := &models.User{
			ID:     1,
			Name:   "Randy Steven",
			Email:  "randy.steven@shopee.com",
			Chance: 1,
		}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		userRepo.On("GetById", mock.Anything, gameReq.UserID).
			Return(user, nil)

		boxesRepo.On("GetNineRandomBoxes", mock.Anything).
			Return(boxes, nil)

		gameRepo.On("Save", mock.Anything, gameReq).
			Return(&games[0], nil)

		res, _ := gameUsecase.PlayGame(ctx, gameReq)

		assert.NotNil(t, res)
	})

	t.Run("should return error while get user by id", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}
		gameReq := &models.Game{
			UserID: 1,
		}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		userRepo.On("GetById", mock.Anything, gameReq.UserID).
			Return(nil, errors.New("mock error"))

		_, err := gameUsecase.PlayGame(ctx, gameReq)

		assert.Error(t, err)
	})

	t.Run("should return failed because user chance is 0", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}
		gameReq := &models.Game{
			UserID: 1,
		}
		user := &models.User{
			ID:     1,
			Name:   "Randy Steven",
			Email:  "randy.steven@shopee.com",
			Chance: 0,
		}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		userRepo.On("GetById", mock.Anything, gameReq.UserID).
			Return(user, nil)

		_, err := gameUsecase.PlayGame(ctx, gameReq)

		assert.Equal(t, &apperror.ErrZeroChance{}, err)
	})

	t.Run("should return failed because failed generate random boxes", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}
		gameReq := &models.Game{
			UserID: 1,
		}
		user := &models.User{
			ID:     1,
			Name:   "Randy Steven",
			Email:  "randy.steven@shopee.com",
			Chance: 1,
		}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		userRepo.On("GetById", mock.Anything, gameReq.UserID).
			Return(user, nil)

		boxesRepo.On("GetNineRandomBoxes", mock.Anything).
			Return(nil, errors.New("mock anything"))

		_, err := gameUsecase.PlayGame(ctx, gameReq)

		assert.Error(t, err)
	})

	t.Run("should return error while create game", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}
		gameReq := &models.Game{
			UserID: 1,
		}
		user := &models.User{
			ID:     1,
			Name:   "Randy Steven",
			Email:  "randy.steven@shopee.com",
			Chance: 1,
		}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		userRepo.On("GetById", mock.Anything, gameReq.UserID).
			Return(user, nil)

		boxesRepo.On("GetNineRandomBoxes", mock.Anything).
			Return(boxes, nil)

		gameRepo.On("Save", mock.Anything, gameReq).
			Return(nil, errors.New("mock error"))

		_, err := gameUsecase.PlayGame(ctx, gameReq)

		assert.Error(t, err)

	})
}

func TestChooseReward(t *testing.T) {
	t.Run("should return success updated reward", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		chooseReward := &req.ChooseReward{
			GameID: 1,
			BoxID:  1,
		}

		gameRepo.On("GetById", mock.Anything, chooseReward.GameID).
			Return(&games[0], nil)

		gameRepo.On("CreateRewardTransaction", mock.Anything, &games[0]).
			Return(&games[0], nil)

		req, _ := gameUsecase.ChooseReward(ctx, chooseReward)

		assert.Equal(t, chooseReward.BoxID, req.WinBoxID)
	})

	t.Run("should return failed due error while query games", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		chooseReward := &req.ChooseReward{
			GameID: 1,
			BoxID:  1,
		}

		gameRepo.On("GetById", mock.Anything, chooseReward.GameID).
			Return(nil, errors.New("mock error"))

		_, err := gameUsecase.ChooseReward(ctx, chooseReward)

		assert.Error(t, err)
	})

	t.Run("should return error due game not found", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		chooseReward := &req.ChooseReward{
			GameID: 1,
			BoxID:  1,
		}

		gameRepo.On("GetById", mock.Anything, chooseReward.GameID).
			Return(nil, nil)

		req, err := gameUsecase.ChooseReward(ctx, chooseReward)

		assert.Error(t, err)
		assert.Nil(t, req)
		assert.Equal(t, &apperror.ErrDataNotFound{Data: "game"}, err)
	})

	t.Run("should return error because win box id already filled", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		chooseReward := &req.ChooseReward{
			GameID: 2,
			BoxID:  1,
		}

		gameRepo.On("GetById", mock.Anything, chooseReward.GameID).
			Return(&games[1], nil)

		_, err := gameUsecase.ChooseReward(ctx, chooseReward)

		assert.Equal(t, &apperror.ErrInvalidRequest{Field: enums.BoxId}, err)
	})

	t.Run("should return error because choosed box is in not games", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		chooseReward := &req.ChooseReward{
			GameID: 2,
			BoxID:  100,
		}

		gameRepo.On("GetById", mock.Anything, chooseReward.GameID).
			Return(&games[1], nil)

		_, err := gameUsecase.ChooseReward(ctx, chooseReward)

		assert.Equal(t, &apperror.ErrInvalidRequest{Field: enums.BoxId}, err)
	})

	t.Run("should return failed while create reward transaction", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		chooseReward := &req.ChooseReward{
			GameID: 1,
			BoxID:  1,
		}

		gameRepo.On("GetById", mock.Anything, chooseReward.GameID).
			Return(&games[0], nil)

		gameRepo.On("CreateRewardTransaction", mock.Anything, &games[0]).
			Return(nil, errors.New("mock error"))

		_, err := gameUsecase.ChooseReward(ctx, chooseReward)

		assert.Error(t, err)
	})
}

func TestGetUserCurrentChance(t *testing.T) {
	t.Run("should return user", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		user := &models.User{
			ID:     1,
			Name:   "Randy Steven",
			Email:  "randy.steven@shopee.com",
			Chance: 1,
		}

		userRepo.On("GetById", mock.Anything, user.ID).
			Return(user, nil)

		userRes, _ := gameUsecase.GetUserCurrentChance(ctx, user.ID)
		assert.Equal(t, user.Chance, userRes.Chance)
	})

	t.Run("should return error", func(t *testing.T) {
		ctx := context.Background()
		gameRepo := mocks.GameRepository{}
		userRepo := mocks.UserRepository{}
		transactionRepo := mocks.TransactionRepository{}
		walletRepo := mocks.WalletRepository{}
		boxesRepo := mocks.BoxRepository{}

		gameUsecase := usecases.NewGameUsecase(
			&gameRepo,
			&userRepo,
			&transactionRepo,
			&walletRepo,
			&boxesRepo,
		)

		userRepo.On("GetById", mock.Anything, uint(0)).
			Return(nil, errors.New("mock error"))

		_, err := gameUsecase.GetUserCurrentChance(ctx, uint(0))
		assert.Error(t, err)
	})
}
