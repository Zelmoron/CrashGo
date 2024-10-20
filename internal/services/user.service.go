package services

import (
	"casego/internal/dto"
	"casego/internal/models"

	"gorm.io/gorm"
)

func CreateUser(userDTO dto.UserDTO, db *gorm.DB) (models.UserModel, error) {
	user := FromDTO(userDTO) // получение модели пользоватля по данным которые пришли в request

	if err := db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetUsers(db *gorm.DB) ([]models.UserModel, error) {
	var users []models.UserModel

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(telegram_id string, db *gorm.DB) (models.UserModel, error) {
	var user models.UserModel

	if err := db.Where("telegram_id=?", telegram_id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(id string, userDTO dto.UserDTO, db *gorm.DB) (models.UserModel, error) {
	var user models.UserModel

	if err := db.First(&user, id).Error; err != nil {
		return user, err
	}

	user.TelegramID = userDTO.Id
	user.Coins = userDTO.Coins

	if err := db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(id string, db *gorm.DB) error {
	var user models.UserModel

	if err := db.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}

func FromDTO(dto dto.UserDTO) models.UserModel {

	//метод для получения данных, которые будут добавляться/проверятся в бд
	return models.UserModel{
		Name:       dto.Name,
		TelegramID: dto.Id,
		Coins:      dto.Coins,
	}
}
