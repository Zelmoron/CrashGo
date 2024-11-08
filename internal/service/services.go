package service

import (
	"CaseGo/internal/models"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB //поле для базы данных
}

func New(db *gorm.DB) *Service {
	//возвращаем указатель на с.Service
	return &Service{

		db: db, // Поле для бд
	}
}

func (s *Service) GetUsers(id int) models.UserModel {
	//бизнес логика для получения пользователя

	var user models.UserModel

	if err := s.db.Where("telegram_id=?", id).First(&user).Error; err != nil {
		return user
	}

	return user
}

func (s *Service) CreateUser() {

}
