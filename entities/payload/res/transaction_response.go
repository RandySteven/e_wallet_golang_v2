package res

import (
	"time"

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
	ReceiverName    string          `json:"receiver_name,omitempty"`
	ReceiverWallet  string          `json:"receiver_wallet,omitempty"`
	TopupUser       string          `json:"topup_user,omitempty"`
	TopupWallet     string          `json:"topup_wallet,omitempty"`
	Description     string          `json:"description"`
	Amount          decimal.Decimal `json:"amount"`
}
