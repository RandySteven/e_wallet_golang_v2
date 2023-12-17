package res

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/entities/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/randy-steven/assignment-go-rest-api/enums"
	"github.com/shopspring/decimal"
)

type TransactionPaginationResponses struct {
	Page         string                      `json:"page,omitempty"`
	Total        uint                        `json:"total,omitempty"`
	Transactions []TransactionDetailResponse `json:"transactions"`
}

type TransactionDetailResponse struct {
	ID              uint            `json:"id"`
	TransactionDate time.Time       `json:"transaction_date"`
	TransactionType string          `json:"transaction_type"`
	SenderName      string          `json:"sender_name,omitempty"`
	SenderWallet    string          `json:"sender_wallet,omitempty"`
	ReceipentName   string          `json:"receipent_name,omitempty"`
	ReceipentWallet string          `json:"receipent_wallet,omitempty"`
	Description     string          `json:"description"`
	Amount          decimal.Decimal `json:"amount"`
}

func TransactionResponse(transaction *models.Transaction) *TransactionDetailResponse {
	transactionDetail := &TransactionDetailResponse{
		ID:              transaction.ID,
		TransactionDate: transaction.CreatedAt,
		Description:     transaction.Description,
		Amount:          transaction.Amount,
	}

	if transaction.ReceiverID == transaction.SenderID {
		transactionDetail.TransactionType = enums.Topup
	} else {
		transactionDetail.TransactionType = enums.Transfer
		transactionDetail.SenderName = transaction.Sender.User.Name
		transactionDetail.SenderWallet = transaction.Sender.Number
	}

	if transaction.Receiver != nil {
		transactionDetail.ReceipentName = transaction.Receiver.User.Name
		transactionDetail.ReceipentWallet = transaction.Receiver.Number
	}

	return transactionDetail
}
