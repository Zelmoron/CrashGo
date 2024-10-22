package models

import (
	"time"

	"gorm.io/gorm"
)

type InventoryModel struct {
	ID         uint           `gorm:"primaryKey"`
	TelegramID uint           `gorm:"not null"`
	Name       string         `gorm:"type:varchar(50)"`
	Type       string         `gorm:"type:varchar(50)"`
	Cost       float32        `gorm:"type:float;default:0" validate:"gte=0"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (InventoryModel) TableName() string {
	return "inventory"
}
