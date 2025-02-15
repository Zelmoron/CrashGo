package models

type ItemModel struct {
	ID         int     `gorm:"primaryKey"`
	WeaponName string  `gorm:"size:255;default:X)"`
	SkinName   string  `gorm:"size:255;default:X)"`
	Image      string  `gorm:"size:500;default:https://cdn1.epicgames.com/offer/9c1a74145a9145ec803d7452e80819a0/EGS_Fallout4_BethesdaGameStudios_S1_2560x1440-fb0e82fa71a74e750c95b57b91fc558d?resize=1&w=480&h=270&quality=medium"`
	Type       string  `gorm:"type:varchar(50)"`
	Price      float64 `gorm:"type:float"`
	CaseID     int
}

func (ItemModel) TableName() string {
	return "items"
}
