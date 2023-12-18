package req

import "github.com/shopspring/decimal"

type TopupRequest struct {
	UserID       uint
	Amount       decimal.Decimal `json:"amount" binding:"required"`
	SourceOfFund uint            `json:"source_of_fund" binding:"required"`
}
