package req

import "github.com/shopspring/decimal"

type TransferRequest struct {
	SenderUserId     uint
	ReceiverWalletId string          `json:"to" binding:"required"`
	Amount           decimal.Decimal `json:"amount" binding:"required"`
	Description      string          `json:"description" binding:"max=35"`
}
