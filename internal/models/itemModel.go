package models

type ItemModel struct {
	ID         int     `gorm:"primaryKey"`
	WeaponName string  `gorm:"size:255"`
	SkinName   string  `gorm:"size:255"`
	Image      string  `gorm:"size:500"`
	Type       string  `gorm:"type:varchar(50)"`
	Price      float64 `gorm:"type:float"`
	CaseID     int
}

func (ItemModel) TableName() string {
	return "items"
}
