package db

import "fmt"

var db = 1

type User struct {
	TelegramId int    `json:"id"`
	Name       string `json:"name"`
	Coins      int
}

func Connect() {
	//здесь подключение к бд и создание таблицы

	// эта функция запускается с мейна один раз, поэтому что сделать так, чтобы переменная для работы с бд была открыта для этого пакета

	db++

}

func (u *User) CreateUser() error {
	//проверка на то, есть ли он в бд
	//если есть, идем дальше, если нет, создаем.

	//идет возврат ошибки

	u.Coins = 100 // типо бонус при создании нового пользователя
	fmt.Println(u)

	return nil

}
