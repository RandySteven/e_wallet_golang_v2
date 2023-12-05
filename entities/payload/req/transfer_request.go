package req

import "github.com/shopspring/decimal"

type TransferRequest struct {
	To          uint            `json:"to"`
	Amount      decimal.Decimal `json:"amount"`
	Description string          `json:"description" binding:"max=35"`
}
