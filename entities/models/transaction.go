package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint            `gorm:"primaryKey;autoIncrement"`
	SenderID       uint            `gorm:"not null"`
	ReceiverID     uint            `gorm:"not null"`
	Amount         decimal.Decimal `gorm:"not null;type:numeric"`
	Description    string          `gorm:"not null"`
	SourceOfFundID uint            `gorm:"not null"`
	CreatedAt      time.Time       `gorm:"not null;default:current_timestamp"`
	UpdatedAt      time.Time       `gorm:"not null;default:current_timestamp"`
	DeletdAt       gorm.DeletedAt

	Sender       *Wallet       `gorm:"foreignKey:SenderID;references:ID"`
	Receiver     *Wallet       `gorm:"foreignKey:ReceiverID;references:ID"`
	SourceOfFund *SourceOfFund `gorm:"foreignKey:SourceOfFundID;references:ID"`
}
