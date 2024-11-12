package service

import (
	"CaseGo/internal/models"
	"fmt"
	"log"

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

func (s *Service) CreateUser(id int, name string) models.UserModel {

	//бизнес логика для проверки/добавления пользователя

	var user models.UserModel

	if err := s.db.Where("telegram_id=?", id).First(&user).Error; err == nil {
		log.Println("Уже есть")
		return user
	}

	user = models.UserModel{
		Name:       name,
		TelegramID: id,
		Coins:      100,
	}
	if err := s.db.Create(&user).Error; err != nil {
		return user
	}
	log.Println("Новый")
	return user

}

type Inventory struct {
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
	Type string  `json:"type"`
}

func (s *Service) GetInventory(id int) []Inventory {
	//бизнес логика доя получения инывентаря
	var inventory []models.InventoryModel

	if err := s.db.Where("telegram_id=?", id).Find(&inventory).Error; err != nil {
		return []Inventory{}
	}

	var inv []Inventory

	for _, v := range inventory {
		inv = append(inv, Inventory{v.Name, v.Cost, v.Type})
	}

	return inv

}

type Cases struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (s *Service) GetCases() []Cases {
	//бизнес логика для получения кейсов

	var cases []models.CasesModel

	if err := s.db.Find(&cases).Error; err != nil {
		return []Cases{}
	}
	var casesAll []Cases
	for _, v := range cases {
		casesAll = append(casesAll, Cases{v.ID, v.Name, v.Image})
	}
	return casesAll
}

type Weapons struct {
	Id      int    `json:"id"`
	Name    string `json:"weapon_name"`
	Skin    string `json:"skin_name"`
	Rarity  string `json:"rarity"`
	Img     string `json:"steam_image"`
	IsLoser bool   `json:"isLoser"`
}

func (s *Service) GetWeapons(id int) []Weapons {
	var cases models.CasesModel

	if err := s.db.Preload("Items").First(&cases, id).Error; err != nil {
		fmt.Println(1)
		return []Weapons{}
	}
	weaponsData := cases.Items

	var weapons []Weapons

	for _, v := range weaponsData {
		weapons = append(weapons, Weapons{Id: v.ID, Name: v.WeaponName, Skin: v.SkinName, Rarity: v.Type, Img: v.Image, IsLoser: false})

	}

	return weapons

}
