package dto

type UserDTO struct {
	//структура, в которую приходит fetch запрос
	Name  string  `json:"name" validate:"required,min=2,max=50"` // имя пользователя
	Id    int     `json:"id" validate:"gte=0"`                   // айди телеграмма!!!!
	Coins float32 `json:"coins" validate:"gte=0"`                // монеты в качестве бонуса
}
