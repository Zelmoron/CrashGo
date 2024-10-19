package models

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(50)"`
	Coins     float32        `gorm:"type:float;default:0" validate:"gte=0"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}
