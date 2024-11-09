package models

type ItemModel struct {
	ID         int    `gorm:"primaryKey"`
	WeaponName string `gorm:"size:255"`
	SkinName   string `gorm:"size:255"`
	Image      string `gorm:"size:255"`
	Type       string `gorm:"type:varchar(50)"`
	CaseID     int
}

func (ItemModel) TableName() string {
	return "items"
}
