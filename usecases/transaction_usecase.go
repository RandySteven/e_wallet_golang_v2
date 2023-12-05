package usecases

import (
	"assignment_4/entities/models"
	"assignment_4/entities/payload/req"
	"assignment_4/interfaces"
	"context"
	"log"
)

type transactionUsecase struct {
	sourceFundRepo  interfaces.SourceOfFundRepository
	walletRepo      interfaces.WalletRepository
	transactionRepo interfaces.TransactionRepository
}

// CreateTopupTransaction implements interfaces.TransactionUsecase.
func (usecase *transactionUsecase) CreateTopupTransaction(ctx context.Context, topup *req.TopupRequest) (*models.Transaction, error) {
	wallet, err := usecase.walletRepo.GetByUserId(ctx, topup.UserID)
	if err != nil {
		return nil, err
	}

	sourceFund, err := usecase.sourceFundRepo.GetSourceOfFundBySource(ctx, topup.SourceOfFund)
	if err != nil {
		return nil, err
	}

	log.Println(wallet)
	log.Println(sourceFund)

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
	return transaction, nil
}

// CreateTransferTransaction implements interfaces.TransactionUsecase.
func (usecase *transactionUsecase) CreateTransferTransaction(ctx context.Context, transfer *req.TransferRequest) (*models.Transaction, error) {

	senderWallet, err := usecase.walletRepo.GetByUserId(ctx, transfer.SenderUserId)
	if err != nil {
		return nil, err
	}

	receiverWallet, err := usecase.walletRepo.GetByNumber(ctx, transfer.ReceiverWalletId)
	if err != nil {
		return nil, err
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
) *transactionUsecase {
	return &transactionUsecase{
		sourceFundRepo:  sourceFundRepo,
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
	}
}

var _ interfaces.TransactionUsecase = &transactionUsecase{}
