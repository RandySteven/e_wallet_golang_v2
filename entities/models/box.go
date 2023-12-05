package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Box struct {
	ID        uint            `gorm:"primaryKey;autoIncrement"`
	Amount    decimal.Decimal `gorm:"not null"`
	CreatedAt time.Time       `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time       `gorm:"not null;default:current_timestamp"`
	DeletedAt gorm.DeletedAt
}
