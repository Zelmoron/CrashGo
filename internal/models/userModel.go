package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	//Стуктура для работы с бд
	ID         uint           `gorm:"primaryKey"`
	TelegramID int            `gorm:"type:integer;default:0" validate:"gte=0"`
	Name       string         `gorm:"type:varchar(50)"`
	Coins      float32        `gorm:"type:float;default:0" validate:"gte=0"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}
