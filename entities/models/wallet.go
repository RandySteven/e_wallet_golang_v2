package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Wallet struct {
	ID        uint            `gorm:"primaryKey;autoIncrement"`
	Number    string          `gorm:"not null;unique;default:concat(100::text, lpad(nextval('my_sequence')::text,10,'0'))"`
	Balance   decimal.Decimal `gorm:"not null"`
	UserID    uint            `gorm:"not null;unique"`
	User      User            `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time       `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time       `gorm:"not null;default:current_timestamp"`
	DeletdAt  gorm.DeletedAt
}
