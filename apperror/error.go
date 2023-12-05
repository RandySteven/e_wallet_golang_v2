package apperror

import (
	"fmt"

	"github.com/shopspring/decimal"
)

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
	MinLength uint
	MaxLength uint
	Field     string
	Err       error
}

type ErrSenderAndReceiverSame struct {
	Message string
	Err     error
}

func (e *ErrSenderAndReceiverSame) Error() string {
	return e.Message
}

func (e *ErrAmountLimit) Error() string {
	return fmt.Sprintf("Amount must between %v and %v", e.Min, e.Max)
}

func (e *ErrLengthValidation) Error() string {
	return fmt.Sprintf("%s length must less than ", e.Field, e.MaxLength)
}
