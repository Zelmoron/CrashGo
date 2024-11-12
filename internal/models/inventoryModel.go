package models

import (
	"time"

	"gorm.io/gorm"
)

type InventoryModel struct {
	ID         uint           `gorm:"primaryKey"`
	TelegramID uint           `gorm:"not null"`
	SkinId     int            `gorm:"size:255"`
	WeaponName string         `gorm:"size:255"`
	SkinName   string         `gorm:"size:255"`
	Image      string         `gorm:"size:500"`
	Type       string         `gorm:"type:varchar(50)"`
	Price      float64        `gorm:"type:float"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (InventoryModel) TableName() string {
	return "inventory"
}
