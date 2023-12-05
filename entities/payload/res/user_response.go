package res

import "github.com/shopspring/decimal"

type UserResponse struct {
	ID    uint
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserDetail struct {
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	WalletNumber string          `json:"wallet_number"`
	Balance      decimal.Decimal `json:"balance"`
}
