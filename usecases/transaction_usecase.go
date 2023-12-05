package usecases

import (
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/interfaces"
	"context"
)

type transactionUsecase struct {
	sourceFundRepo interfaces.SourceOfFundRepository
	walletRepo     interfaces.WalletRepository
}

// CreateTopupTransaction implements interfaces.TransactionUsecase.
func (usecase *transactionUsecase) CreateTopupTransaction(ctx context.Context, topup *req.TopupRequest) (*models.Transaction, error) {
	return nil, nil
}

// CreateTransferTransaction implements interfaces.TransactionUsecase.
func (usecase *transactionUsecase) CreateTransferTransaction(ctx context.Context, transfer *req.TransferRequest) (*models.Transaction, error) {
	// receiverWallet, err := usecase.walletRepo.GetByNumber(ctx, transfer.Receiver)
	// if err != nil{
	// 	return nil, err
	// }
	return nil, nil
}

func NewTransactionUsecase(
	sourceFundRepo interfaces.SourceOfFundRepository,
	walletRepo interfaces.WalletRepository,
) *transactionUsecase {
	return &transactionUsecase{
		sourceFundRepo: sourceFundRepo,
		walletRepo:     walletRepo,
	}
}

var _ interfaces.TransactionUsecase = &transactionUsecase{}
