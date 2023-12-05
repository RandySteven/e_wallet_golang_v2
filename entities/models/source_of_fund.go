package models

import (
	"time"

	"gorm.io/gorm"
)

type SourceOfFund struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Source    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	DeletdAt  gorm.DeletedAt
}
