package service

import (
	"CaseGo/internal/models"
)

type Repository interface {
	InsertUser(int, string) models.UserModel
	SelectUser(int) models.UserModel
	SelectInventory(int) []models.InventoryModel
	SelectCases() []models.CasesModel
	SelectWeapons(int) []models.ItemModel
	InsertInventory(int, int) models.ItemModel
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	//возвращаем указатель на с.Service
	return &Service{

		repository: repository,
	}
}

func (s *Service) GetUsers(id int) models.UserModel {
	//бизнес логика для получения пользователя

	user := s.repository.SelectUser(id)

	return user
}

func (s *Service) CreateUser(id int, name string) models.UserModel {

	//бизнес логика для проверки/добавления пользователя

	user := s.repository.InsertUser(id, name)

	return user

}

type Inventory struct {
	Id     int    `json:"id"`
	Name   string `json:"weapon_name"`
	Skin   string `json:"skin_name"`
	Rarity string `json:"rarity"`
	Img    string `json:"steam_image"`
}

func (s *Service) GetInventory(id int) []Inventory {
	//бизнес логика доя получения инывентаря

	inventory := s.repository.SelectInventory(id)
	var inv []Inventory

	for _, v := range inventory {
		inv = append(inv, Inventory{v.SkinId, v.WeaponName, v.SkinName, v.Type, v.Image})
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

	cases := s.repository.SelectCases()
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

	weaponsData := s.repository.SelectWeapons(id)
	var weapons []Weapons

	for _, v := range weaponsData {
		weapons = append(weapons, Weapons{Id: v.ID, Name: v.WeaponName, Skin: v.SkinName, Rarity: v.Type, Img: v.Image, IsLoser: false})

	}

	return weapons

}

func (s *Service) OpenCase(userId int, itemId int) models.ItemModel {
	item := s.repository.InsertInventory(userId, itemId)

	return item

}
