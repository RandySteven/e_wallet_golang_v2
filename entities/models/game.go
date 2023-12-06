package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	UserID   uint `gorm:"not null"`
	BoxID1   uint `gorm:"not null"`
	BoxID2   uint `gorm:"not null"`
	BoxID3   uint `gorm:"not null"`
	BoxID4   uint `gorm:"not null"`
	BoxID5   uint `gorm:"not null"`
	BoxID6   uint `gorm:"not null"`
	BoxID7   uint `gorm:"not null"`
	BoxID8   uint `gorm:"not null"`
	BoxID9   uint `gorm:"not null"`
	WinBoxID uint

	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt gorm.DeletedAt

	Box1 *Box `gorm:"foreignKey:BoxID1;references:ID" json:"omitempty"`
	Box2 *Box `gorm:"foreignKey:BoxID2;references:ID" json:"omitempty"`
	Box3 *Box `gorm:"foreignKey:BoxID3;references:ID" json:"omitempty"`
	Box4 *Box `gorm:"foreignKey:BoxID4;references:ID" json:"omitempty"`
	Box5 *Box `gorm:"foreignKey:BoxID5;references:ID" json:"omitempty"`
	Box6 *Box `gorm:"foreignKey:BoxID6;references:ID" json:"omitempty"`
	Box7 *Box `gorm:"foreignKey:BoxID7;references:ID" json:"omitempty"`
	Box8 *Box `gorm:"foreignKey:BoxID8;references:ID" json:"omitempty"`
	Box9 *Box `gorm:"foreignKey:BoxID9;references:ID" json:"omitempty"`

	WinBox *Box `gorm:"foreignKey:WinBoxID;references:ID" json:"omitempty"`
}
