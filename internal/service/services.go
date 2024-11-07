package service

import (
	"CaseGo/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetUsers(database *gorm.DB) {
	var user models.UserModel

	if err := database.Where("telegram_id=?", 1628918728).First(&user).Error; err != nil {
		fmt.Println("1")
	}

	fmt.Println(user)
}
