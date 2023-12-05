package req

import "github.com/shopspring/decimal"

type TopupRequest struct {
	Amount       decimal.Decimal `json:"amount"`
	SourceOfFund string          `json:"source_of_fund"`
}
