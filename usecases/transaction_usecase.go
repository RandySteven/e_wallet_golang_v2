package usecases

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/req"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/payload/res"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/interfaces"

	"github.com/shopspring/decimal"
)

type transactionUsecase struct {
	sourceFundRepo  interfaces.SourceOfFundRepository
	walletRepo      interfaces.WalletRepository
	transactionRepo interfaces.TransactionRepository
	userRepo        interfaces.UserRepository
}

// GetAllTransactionsRecords implements interfaces.TransactionUsecase.
func (usecase *transactionUsecase) GetAllTransactionsRecords(ctx context.Context, query *entities.QueryCondition, userId uint) (*res.TransactionPaginationResponses, error) {
	wallet, err := usecase.walletRepo.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	transactions, err := usecase.transactionRepo.GetAllTransactions(ctx, query, wallet.ID)
	if err != nil {
		return nil, err
	}
	totalItems, err := usecase.transactionRepo.GetTransactionCountBasedUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	var transactionDetails = []res.TransactionDetailResponse{}
	for _, transaction := range transactions {
		transactionDetail := res.TransactionResponse(&transaction)

		transactionDetails = append(transactionDetails, *transactionDetail)
	}

	transactionPage := &res.TransactionPaginationResponses{
		Page:         query.Page,
		Total:        totalItems,
		Transactions: transactionDetails,
	}
	return transactionPage, nil
}

// CreateTopupTransaction implements interfaces.TransactionUsecase.
func (usecase *transactionUsecase) CreateTopupTransaction(ctx context.Context, topup *req.TopupRequest) (*models.Transaction, error) {
	wallet, err := usecase.walletRepo.GetByUserId(ctx, topup.UserID)
	if err != nil {
		return nil, err
	}

	if topup.SourceOfFund == enums.Reward {
		return nil, &apperror.ErrInvalidRequest{Field: enums.SourceOfFund}
	}

	sourceFund, err := usecase.sourceFundRepo.GetById(ctx, topup.SourceOfFund)
	if err != nil {
		return nil, err
	}

	if sourceFund == nil {
		return nil, &apperror.ErrInvalidRequest{Field: enums.SourceOfFund}
	}

	if decimal.Min(topup.Amount, decimal.NewFromInt(enums.MIN_TOPUP_AMOUNT-1)) == topup.Amount ||
		decimal.Max(topup.Amount, decimal.NewFromInt(enums.MAX_TOPUP_AMOUNT+1)) == topup.Amount {
		return nil, &apperror.ErrAmountLimit{
			Min: decimal.NewFromInt(enums.MIN_TOPUP_AMOUNT),
			Max: decimal.NewFromInt(enums.MAX_TOPUP_AMOUNT),
		}
	}

	transaction := &models.Transaction{
		SenderID:       wallet.ID,
		ReceiverID:     wallet.ID,
		Amount:         topup.Amount,
		SourceOfFundID: sourceFund.ID,
		Description:    "Top up from " + sourceFund.Source,
	}

	transaction, err = usecase.transactionRepo.CreateTopupTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	if transaction.Amount.Equal(decimal.NewFromInt(enums.MAX_TOPUP_AMOUNT)) {
		user, err := usecase.userRepo.GetById(ctx, topup.UserID)
		if err != nil {
			return nil, err
		}

		user.Chance += 1
		_, err = usecase.userRepo.Update(ctx, user)
		if err != nil {
			return nil, err
		}
	}

	return transaction, nil
}

// CreateTransferTransaction implements interfaces.TransactionUsecase.
func (usecase *transactionUsecase) CreateTransferTransaction(ctx context.Context, transfer *req.TransferRequest) (*models.Transaction, error) {

	senderWallet, err := usecase.walletRepo.GetByUserId(ctx, transfer.SenderUserId)
	if err != nil {
		return nil, err
	}

	if senderWallet.Balance.Cmp(transfer.Amount) == -1 {
		return nil, &apperror.ErrBalanceNotEnough{}
	}

	receiverWallet, err := usecase.walletRepo.GetByNumber(ctx, transfer.ReceiverWalletId)
	if err != nil {
		return nil, err
	}

	if receiverWallet == nil {
		return nil, &apperror.ErrWalletInvalid{}
	}

	if senderWallet.ID == receiverWallet.ID {
		return nil, &apperror.ErrSenderAndReceiverSame{
			Message: "user can't transfer money to theirself",
		}
	}

	if decimal.Min(transfer.Amount, decimal.NewFromInt(enums.MIN_TRANSFER_AMOUNT-1)) == transfer.Amount ||
		decimal.Max(transfer.Amount, decimal.NewFromInt(enums.MAX_TRANSFER_AMOUNT+1)) == transfer.Amount {
		return nil, &apperror.ErrAmountLimit{
			Min: decimal.NewFromInt(enums.MIN_TRANSFER_AMOUNT),
			Max: decimal.NewFromInt(enums.MAX_TRANSFER_AMOUNT),
		}
	}

	transaction := &models.Transaction{
		SenderID:       senderWallet.ID,
		ReceiverID:     receiverWallet.ID,
		Amount:         transfer.Amount,
		Description:    transfer.Description,
		SourceOfFundID: 5,
	}

	transaction, err = usecase.transactionRepo.CreateTransferTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func NewTransactionUsecase(
	sourceFundRepo interfaces.SourceOfFundRepository,
	walletRepo interfaces.WalletRepository,
	transactionRepo interfaces.TransactionRepository,
	userRepo interfaces.UserRepository,
) *transactionUsecase {
	return &transactionUsecase{
		sourceFundRepo:  sourceFundRepo,
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
	}
}

var _ interfaces.TransactionUsecase = &transactionUsecase{}
