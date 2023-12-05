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

	Box1 *Box `gorm:"foreignKey:BoxID1;references:ID"`
	Box2 *Box `gorm:"foreignKey:BoxID2;references:ID"`
	Box3 *Box `gorm:"foreignKey:BoxID3;references:ID"`
	Box4 *Box `gorm:"foreignKey:BoxID4;references:ID"`
	Box5 *Box `gorm:"foreignKey:BoxID5;references:ID"`
	Box6 *Box `gorm:"foreignKey:BoxID6;references:ID"`
	Box7 *Box `gorm:"foreignKey:BoxID7;references:ID"`
	Box8 *Box `gorm:"foreignKey:BoxID8;references:ID"`
	Box9 *Box `gorm:"foreignKey:BoxID9;references:ID"`

	WinBox *Box `gorm:"foreignKey:WinBoxID;references:ID"`
}
