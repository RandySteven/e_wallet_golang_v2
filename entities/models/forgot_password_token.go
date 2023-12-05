package models

import (
	"time"

	"gorm.io/gorm"
)

type ForgotPasswordToken struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	ResetToken  string `gorm:"not null"`
	TokenExpiry time.Time
	IsValid     bool `gorm:"not null"`
	UserID      uint `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt gorm.DeletedAt

	User User `gorm:"foreignKey:UserID;references:ID"`
}
