package res

import "github.com/shopspring/decimal"

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type UserLoginResponse struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Token string `json:"token,omitempty"`
}

type UserDetail struct {
	ID           uint            `json:"id"`
	Name         string          `json:"name,omitempty"`
	Email        string          `json:"email,omitempty"`
	WalletNumber string          `json:"wallet_number,omitempty"`
	Balance      decimal.Decimal `json:"balance,omitempty"`
	Chance       uint            `json:"chance,omitempty"`
}
