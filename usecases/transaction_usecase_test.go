package usecases_test

import (
	"assignment_4/apperror"
	"assignment_4/entities"
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

var transactions = []models.Transaction{
	{
		ID:             1,
		SenderID:       1,
		ReceiverID:     2,
		Amount:         decimal.NewFromInt(50000),
		Description:    "",
		SourceOfFundID: 5,
		Sender: &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(60000),
			UserID:  1,
			User: &models.User{
				ID:    1,
				Name:  "Randy Steven",
				Email: "randy.steven@gmail.com",
			},
		},
		Receiver: &models.Wallet{
			ID:      2,
			Number:  "10000000000002",
			Balance: decimal.NewFromInt(60000),
			UserID:  2,
			User: &models.User{
				ID:    2,
				Name:  "Matthew Alfredo",
				Email: "matthew.alfredo@gmail.com",
			},
		},
	},
	{
		ID:             2,
		SenderID:       1,
		ReceiverID:     1,
		Amount:         decimal.NewFromInt(50000),
		Description:    "",
		SourceOfFundID: 1,
		Sender: &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(60000),
			UserID:  1,
			User: &models.User{
				ID:    1,
				Name:  "Randy Steven",
				Email: "randy.steven@gmail.com",
			},
		},
		Receiver: &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(60000),
			UserID:  1,
			User: &models.User{
				ID:    1,
				Name:  "Randy Steven",
				Email: "randy.steven@gmail.com",
			},
		},
	},
}

func TestCreateTransferTransaction(t *testing.T) {
	t.Run("should return success create transfer transaction", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		transferReq := &req.TransferRequest{
			SenderUserId:     1,
			ReceiverWalletId: "10000000000002",
			Amount:           decimal.NewFromInt(50000),
			Description:      "Here you are",
		}

		senderWallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}

		receiverWallet := &models.Wallet{
			ID:      2,
			Number:  "10000000000002",
			Balance: decimal.NewFromInt(1000000),
		}

		transaction := &models.Transaction{
			SenderID:       1,
			ReceiverID:     2,
			Amount:         transferReq.Amount,
			Description:    transferReq.Description,
			SourceOfFundID: 5,
		}

		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, transferReq.SenderUserId).
			Return(senderWallet, nil)

		walletRepo.On("GetByNumber", mock.Anything, transferReq.ReceiverWalletId).
			Return(receiverWallet, nil)

		transactionRepo.On("CreateTransferTransaction", mock.Anything, mock.AnythingOfType("*models.Transaction")).
			Return(transaction, nil)

		ctx := context.Background()
		result, _ := usecase.CreateTransferTransaction(ctx, transferReq)

		assert.Equal(t, transaction, result)

	})

	t.Run("should return error while try to get senderWallet by user id", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		transferReq := &req.TransferRequest{
			SenderUserId:     1,
			ReceiverWalletId: "10000000000002",
			Amount:           decimal.NewFromInt(50000),
			Description:      "Here you are",
		}

		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, transferReq.SenderUserId).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := usecase.CreateTransferTransaction(ctx, transferReq)

		assert.Error(t, err)
	})

	t.Run("should return error while try to get receiverWallet by wallet number", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		transferReq := &req.TransferRequest{
			SenderUserId:     1,
			ReceiverWalletId: "10000000000002",
			Amount:           decimal.NewFromInt(50000),
			Description:      "Here you are",
		}

		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		senderWallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}

		walletRepo.On("GetByUserId", mock.Anything, transferReq.SenderUserId).
			Return(senderWallet, nil)

		walletRepo.On("GetByNumber", mock.Anything, transferReq.ReceiverWalletId).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := usecase.CreateTransferTransaction(ctx, transferReq)

		assert.Error(t, err)
	})

	t.Run("should return failed while create transfer transaction", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		transferReq := &req.TransferRequest{
			SenderUserId:     1,
			ReceiverWalletId: "10000000000002",
			Amount:           decimal.NewFromInt(50000),
			Description:      "Here you are",
		}

		senderWallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}

		receiverWallet := &models.Wallet{
			ID:      2,
			Number:  "10000000000002",
			Balance: decimal.NewFromInt(1000000),
		}

		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, transferReq.SenderUserId).
			Return(senderWallet, nil)

		walletRepo.On("GetByNumber", mock.Anything, transferReq.ReceiverWalletId).
			Return(receiverWallet, nil)

		transactionRepo.On("CreateTransferTransaction", mock.Anything, mock.AnythingOfType("*models.Transaction")).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := usecase.CreateTransferTransaction(ctx, transferReq)

		assert.Error(t, err)
	})

	t.Run("should return error while try transfer amount more than sender balance", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		transferReq := &req.TransferRequest{
			SenderUserId:     1,
			ReceiverWalletId: "10000000000002",
			Amount:           decimal.NewFromInt(50000),
			Description:      "Here you are",
		}

		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		senderWallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000),
		}

		walletRepo.On("GetByUserId", mock.Anything, transferReq.SenderUserId).
			Return(senderWallet, nil)

		ctx := context.Background()
		_, err := usecase.CreateTransferTransaction(ctx, transferReq)

		assert.Error(t, err)
	})

	t.Run("should return failed while user try to transfer to his/her own wallet", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		transferReq := &req.TransferRequest{
			SenderUserId:     1,
			ReceiverWalletId: "10000000000001",
			Amount:           decimal.NewFromInt(50000),
			Description:      "Here you are",
		}

		senderWallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}

		receiverWallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}

		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, transferReq.SenderUserId).
			Return(senderWallet, nil)

		walletRepo.On("GetByNumber", mock.Anything, transferReq.ReceiverWalletId).
			Return(receiverWallet, nil)

		ctx := context.Background()
		_, err := usecase.CreateTransferTransaction(ctx, transferReq)

		assert.Error(t, err)
	})

	t.Run("should return failed while receiver wallet is nil", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		transferReq := &req.TransferRequest{
			SenderUserId:     1,
			ReceiverWalletId: "10000000000001",
			Amount:           decimal.NewFromInt(50000),
			Description:      "Here you are",
		}

		senderWallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}

		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, transferReq.SenderUserId).
			Return(senderWallet, nil)

		walletRepo.On("GetByNumber", mock.Anything, transferReq.ReceiverWalletId).
			Return(nil, nil)

		ctx := context.Background()
		_, err := usecase.CreateTransferTransaction(ctx, transferReq)

		assert.Error(t, err)
	})

	t.Run("should return failed while transfer amount < 1000", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		transferReq := &req.TransferRequest{
			SenderUserId:     1,
			ReceiverWalletId: "10000000000001",
			Amount:           decimal.NewFromInt(999),
			Description:      "Here you are",
		}

		senderWallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}

		receiverWallet := &models.Wallet{
			ID:      2,
			Number:  "10000000000002",
			Balance: decimal.NewFromInt(1000000),
		}

		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, transferReq.SenderUserId).
			Return(senderWallet, nil)

		walletRepo.On("GetByNumber", mock.Anything, transferReq.ReceiverWalletId).
			Return(receiverWallet, nil)

		ctx := context.Background()
		_, err := usecase.CreateTransferTransaction(ctx, transferReq)

		assert.Error(t, err)
	})

	t.Run("should return failed while transfer amount > 50000000", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		transferReq := &req.TransferRequest{
			SenderUserId:     1,
			ReceiverWalletId: "10000000000001",
			Amount:           decimal.NewFromInt(50000001),
			Description:      "Here you are",
		}

		senderWallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}

		receiverWallet := &models.Wallet{
			ID:      2,
			Number:  "10000000000002",
			Balance: decimal.NewFromInt(1000000),
		}

		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, transferReq.SenderUserId).
			Return(senderWallet, nil)

		walletRepo.On("GetByNumber", mock.Anything, transferReq.ReceiverWalletId).
			Return(receiverWallet, nil)

		ctx := context.Background()
		_, err := usecase.CreateTransferTransaction(ctx, transferReq)

		assert.Error(t, err)
	})

}

func TestCreateTopupTransaction(t *testing.T) {
	t.Run("should return success create transaction for top up and success update user chance", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		topupReq := &req.TopupRequest{
			UserID:       1,
			Amount:       decimal.NewFromInt(enums.MAX_TOPUP_AMOUNT),
			SourceOfFund: "Bank Transfer",
		}
		sourceOfFund := &models.SourceOfFund{
			ID:     1,
			Source: "Bank Transfer",
		}
		transaction := &models.Transaction{
			SenderID:       1,
			ReceiverID:     2,
			Amount:         topupReq.Amount,
			SourceOfFundID: 5,
		}
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}
		user := &models.User{
			ID:     1,
			Name:   "Randy Steven",
			Email:  "randy.steven@gmail.com",
			Chance: 0,
		}
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, topupReq.UserID).
			Return(wallet, nil)

		sourceFundRepo.On("GetSourceOfFundBySource", mock.Anything, topupReq.SourceOfFund).
			Return(sourceOfFund, nil)

		transactionRepo.
			On("CreateTopupTransaction", mock.Anything, mock.AnythingOfType("*models.Transaction")).
			Return(transaction, nil)

		userRepo.On("GetById", mock.Anything, topupReq.UserID).
			Return(user, nil)

		userRepo.On("Update", mock.Anything, user).
			Return(user, nil)

		ctx := context.Background()
		res, _ := usecase.CreateTopupTransaction(ctx, topupReq)

		assert.Equal(t, transaction.Amount, res.Amount)
		assert.Equal(t, uint(1), user.Chance)
	})

	t.Run("should return success create transaction for top up and success update user chance", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		topupReq := &req.TopupRequest{
			UserID:       1,
			Amount:       decimal.NewFromInt(50000),
			SourceOfFund: "Bank Transfer",
		}
		sourceOfFund := &models.SourceOfFund{
			ID:     1,
			Source: "Bank Transfer",
		}
		transaction := &models.Transaction{
			SenderID:       1,
			ReceiverID:     2,
			Amount:         topupReq.Amount,
			SourceOfFundID: 5,
		}
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, topupReq.UserID).
			Return(wallet, nil)

		sourceFundRepo.On("GetSourceOfFundBySource", mock.Anything, topupReq.SourceOfFund).
			Return(sourceOfFund, nil)

		transactionRepo.
			On("CreateTopupTransaction", mock.Anything, mock.AnythingOfType("*models.Transaction")).
			Return(transaction, nil)

		ctx := context.Background()
		res, _ := usecase.CreateTopupTransaction(ctx, topupReq)

		assert.Equal(t, transaction.Amount, res.Amount)
	})

	t.Run("should return failed while get wallet by user id error", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		topupReq := &req.TopupRequest{
			UserID:       1,
			Amount:       decimal.NewFromInt(50000),
			SourceOfFund: "Bank Transfer",
		}
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)
		ctx := context.Background()

		walletRepo.On("GetByUserId", mock.Anything, topupReq.UserID).
			Return(nil, errors.New("mock error"))

		_, err := usecase.CreateTopupTransaction(ctx, topupReq)

		assert.Error(t, err)
	})

	t.Run("should return failed while source of fund from rewards", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		topupReq := &req.TopupRequest{
			UserID:       1,
			Amount:       decimal.NewFromInt(50000),
			SourceOfFund: "Reward",
		}
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)
		ctx := context.Background()
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}

		walletRepo.On("GetByUserId", mock.Anything, topupReq.UserID).
			Return(wallet, nil)

		_, err := usecase.CreateTopupTransaction(ctx, topupReq)

		assert.Error(t, err)
		assert.Equal(t, &apperror.ErrInvalidRequest{Field: enums.SourceOfFund}, err)
	})

	t.Run("should return failed while try to query source of fund", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		topupReq := &req.TopupRequest{
			UserID:       1,
			Amount:       decimal.NewFromInt(50000),
			SourceOfFund: "Bank Transfer",
		}
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, topupReq.UserID).
			Return(wallet, nil)

		sourceFundRepo.On("GetSourceOfFundBySource", mock.Anything, topupReq.SourceOfFund).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := usecase.CreateTopupTransaction(ctx, topupReq)

		assert.Error(t, err)
	})

	t.Run("should return failed while try to query source of fund not found", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		topupReq := &req.TopupRequest{
			UserID:       1,
			Amount:       decimal.NewFromInt(50000),
			SourceOfFund: "Lalal",
		}
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, topupReq.UserID).
			Return(wallet, nil)

		sourceFundRepo.On("GetSourceOfFundBySource", mock.Anything, topupReq.SourceOfFund).
			Return(nil, nil)

		ctx := context.Background()
		_, err := usecase.CreateTopupTransaction(ctx, topupReq)

		assert.Error(t, err)
		assert.Equal(t, &apperror.ErrInvalidRequest{Field: enums.SourceOfFund}, err)
	})

	t.Run("should return failed while error create transaction for top up", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		topupReq := &req.TopupRequest{
			UserID:       1,
			Amount:       decimal.NewFromInt(50000),
			SourceOfFund: "Bank Transfer",
		}
		sourceOfFund := &models.SourceOfFund{
			ID:     1,
			Source: "Bank Transfer",
		}
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, topupReq.UserID).
			Return(wallet, nil)

		sourceFundRepo.On("GetSourceOfFundBySource", mock.Anything, topupReq.SourceOfFund).
			Return(sourceOfFund, nil)

		transactionRepo.
			On("CreateTopupTransaction", mock.Anything, mock.AnythingOfType("*models.Transaction")).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := usecase.CreateTopupTransaction(ctx, topupReq)

		assert.Error(t, err)
	})

	t.Run("should return error while try to get user id for update chance", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		topupReq := &req.TopupRequest{
			UserID:       1,
			Amount:       decimal.NewFromInt(enums.MAX_TRANSFER_AMOUNT),
			SourceOfFund: "Bank Transfer",
		}
		sourceOfFund := &models.SourceOfFund{
			ID:     1,
			Source: "Bank Transfer",
		}
		transaction := &models.Transaction{
			SenderID:       1,
			ReceiverID:     1,
			Amount:         topupReq.Amount,
			SourceOfFundID: 5,
		}
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, topupReq.UserID).
			Return(wallet, nil)

		sourceFundRepo.On("GetSourceOfFundBySource", mock.Anything, topupReq.SourceOfFund).
			Return(sourceOfFund, nil)

		transactionRepo.
			On("CreateTopupTransaction", mock.Anything, mock.AnythingOfType("*models.Transaction")).
			Return(transaction, nil)

		userRepo.On("GetById", mock.Anything, topupReq.UserID).
			Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := usecase.CreateTopupTransaction(ctx, topupReq)

		assert.Error(t, err)
	})

	t.Run("should return error while try to update user chance", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		topupReq := &req.TopupRequest{
			UserID:       1,
			Amount:       decimal.NewFromInt(enums.MAX_TRANSFER_AMOUNT),
			SourceOfFund: "Bank Transfer",
		}
		sourceOfFund := &models.SourceOfFund{
			ID:     1,
			Source: "Bank Transfer",
		}
		transaction := &models.Transaction{
			SenderID:       1,
			ReceiverID:     1,
			Amount:         topupReq.Amount,
			SourceOfFundID: 5,
		}
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
		}
		user := &models.User{
			ID:     1,
			Name:   "Randy Steven",
			Email:  "randy.steven@gmail.com",
			Chance: 0,
		}
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)

		walletRepo.On("GetByUserId", mock.Anything, topupReq.UserID).
			Return(wallet, nil)

		sourceFundRepo.On("GetSourceOfFundBySource", mock.Anything, topupReq.SourceOfFund).
			Return(sourceOfFund, nil)

		transactionRepo.
			On("CreateTopupTransaction", mock.Anything, mock.AnythingOfType("*models.Transaction")).
			Return(transaction, nil)

		userRepo.On("GetById", mock.Anything, topupReq.UserID).
			Return(user, nil)

		userRepo.On("Update", mock.Anything, user).Return(nil, errors.New("mock error"))

		ctx := context.Background()
		_, err := usecase.CreateTopupTransaction(ctx, topupReq)

		assert.Error(t, err)
	})
}

func TestGetAllTransactionsRecords(t *testing.T) {
	t.Run("should return res transaction pagination", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		ctx := context.Background()
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)
		count := 0
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
			User: &models.User{
				ID:     1,
				Name:   "Randy Steven",
				Email:  "randy.steven@gmail.com",
				Chance: 0,
			},
		}
		query := &entities.QueryCondition{
			Page: "1",
		}

		walletRepo.On("GetByUserId", mock.Anything, uint(1)).
			Return(wallet, nil)

		transactionRepo.On("GetAllTransactions",
			mock.Anything,
			mock.AnythingOfType("*entities.QueryCondition"),
			uint(1),
		).
			Return(transactions, nil)

		transactionRepo.On("GetTransactionCountBasedUserId",
			mock.Anything,
			uint(1),
		).Return(uint(count), nil)

		res, _ := usecase.GetAllTransactionsRecords(ctx, query, 1)

		assert.Equal(t, query.Page, res.Page)
	})

	t.Run("should return error while try to get user wallet by user id", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		ctx := context.Background()
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)
		query := &entities.QueryCondition{
			Page: "1",
		}

		walletRepo.On("GetByUserId", mock.Anything, uint(1)).
			Return(nil, errors.New("mock error"))

		_, err := usecase.GetAllTransactionsRecords(ctx, query, 1)

		assert.Error(t, err)
	})

	t.Run("should return error while get transactions", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		ctx := context.Background()
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
			User: &models.User{
				ID:     1,
				Name:   "Randy Steven",
				Email:  "randy.steven@gmail.com",
				Chance: 0,
			},
		}
		query := &entities.QueryCondition{
			Page: "1",
		}

		walletRepo.On("GetByUserId", mock.Anything, uint(1)).
			Return(wallet, nil)

		transactionRepo.On("GetAllTransactions",
			mock.Anything,
			mock.AnythingOfType("*entities.QueryCondition"),
			uint(1),
		).
			Return(nil, errors.New("mock error"))

		_, err := usecase.GetAllTransactionsRecords(ctx, query, 1)

		assert.Error(t, err)
	})

	t.Run("should return error while try to get count total transactions", func(t *testing.T) {
		sourceFundRepo := mocks.SourceOfFundRepository{}
		walletRepo := mocks.WalletRepository{}
		transactionRepo := mocks.TransactionRepository{}
		userRepo := mocks.UserRepository{}
		ctx := context.Background()
		usecase := usecases.NewTransactionUsecase(
			&sourceFundRepo,
			&walletRepo,
			&transactionRepo,
			&userRepo,
		)
		wallet := &models.Wallet{
			ID:      1,
			Number:  "10000000000001",
			Balance: decimal.NewFromInt(1000000),
			User: &models.User{
				ID:     1,
				Name:   "Randy Steven",
				Email:  "randy.steven@gmail.com",
				Chance: 0,
			},
		}
		query := &entities.QueryCondition{
			Page: "1",
		}

		walletRepo.On("GetByUserId", mock.Anything, uint(1)).
			Return(wallet, nil)

		transactionRepo.On("GetAllTransactions",
			mock.Anything,
			mock.AnythingOfType("*entities.QueryCondition"),
			uint(1),
		).
			Return(transactions, nil)

		transactionRepo.On("GetTransactionCountBasedUserId",
			mock.Anything,
			uint(1),
		).Return(uint(0), errors.New("mock error"))

		_, err := usecase.GetAllTransactionsRecords(ctx, query, 1)

		assert.Error(t, err)
	})
}
