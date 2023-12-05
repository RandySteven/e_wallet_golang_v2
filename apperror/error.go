package apperror

import "github.com/shopspring/decimal"

type ErrWalletNumberInvalid struct {
	Message string
	Err     error
}

type ErrAmountInvalid struct {
	Message string
	Err     error
}

type ErrFieldIsRequired struct {
	Field   string
	Message string
	Err     error
}

type ErrAmountLimit struct {
	Min     decimal.Decimal
	Max     decimal.Decimal
	Message string
	Err     error
}

type ErrLengthValidation struct {
	Length  uint
	Message string
	Err     error
}
