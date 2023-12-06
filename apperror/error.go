package apperror

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type ErrWalletNumberInvalid struct {
	Message string
	Err     error
}

type ErrInvalidRequest struct {
	Field string
	Err   error
}

type ErrAmountInvalid struct {
	Message string
	Err     error
}

type ErrFieldValidation struct {
	Message []string
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

type ErrDataNotFound struct {
	Data string
	Err  error
}

type ErrWalletInvalid struct {
	Err error
}

type ErrBalanceNotEnough struct {
	Err error
}

type ErrZeroChance struct {
	Err error
}

type ErrTokenExpired struct {
	Err error
}

type ErrTokenInvalid struct {
	Err error
}

type ErrTokenAlreadyUsed struct {
	Err error
}

type ErrInvalidFormat struct {
	Message string
	Err     error
}

type ErrLogin struct {
	Err error
}

type ErrEmailAlreadyExists struct {
	Err error
}

func (e *ErrSenderAndReceiverSame) Error() string {
	return e.Message
}

func (e *ErrAmountLimit) Error() string {
	return fmt.Sprintf("Amount must between %v and %v", e.Min, e.Max)
}

func (e *ErrLengthValidation) Error() string {
	return fmt.Sprintf("%s length must less than %d", e.Field, e.MaxLength)
}

func (e *ErrDataNotFound) Error() string {
	return fmt.Sprintf("%s not found", e.Data)
}

func (e *ErrWalletInvalid) Error() string {
	return "wallet is invalid"
}

func (e *ErrBalanceNotEnough) Error() string {
	return "user don't have enough balance"
}

func (e *ErrInvalidRequest) Error() string {
	return fmt.Sprintf("%s invalid request", e.Field)
}

func (e *ErrZeroChance) Error() string {
	return "the chance user has is 0"
}

func (e *ErrTokenExpired) Error() string {
	return "token already expired"
}

func (e *ErrTokenAlreadyUsed) Error() string {
	return "token is already used"
}

func (e *ErrTokenInvalid) Error() string {
	return "token is invalid"
}

func (e *ErrFieldValidation) Error() string {
	return strings.Join(e.Message, "|")
}

func (e *ErrInvalidFormat) Error() string {
	return e.Message
}

func (e *ErrLogin) Error() string {
	return "email and password didn't match in user"
}

func (e *ErrEmailAlreadyExists) Error() string {
	return "email already exists"
}
