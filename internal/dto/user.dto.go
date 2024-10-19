package dto

type UserDTO struct {
    Name  string  `json:"name" validate:"required,min=2,max=50"`
    Coins float32 `json:"coins" validate:"gte=0"`
}
