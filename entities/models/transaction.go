package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	SenderID       uint
	ReceiverID     uint
	Amount         decimal.Decimal `gorm:"not null"`
	Description    string          `gorm:"not null"`
	SourceOfFundID uint
	CreatedAt      time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt      time.Time `gorm:"not null;default:current_timestamp"`
	DeletdAt       gorm.DeletedAt

	Sender       Wallet       `gorm:"foreignKey:SenderID;references:ID"`
	Receiver     Wallet       `gorm:"foreignKey:ReceiverID;references:ID"`
	SourceOfFund SourceOfFund `gorm:"foreignKey:SourceOfFundID;references:ID"`
}
