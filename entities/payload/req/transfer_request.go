package req

import "github.com/shopspring/decimal"

type TransferRequest struct {
	Sender      string
	Receiver    string          `json:"to"`
	Amount      decimal.Decimal `json:"amount"`
	Description string          `json:"description" binding:"max=35"`
}
