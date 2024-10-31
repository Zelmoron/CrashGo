package models

type Cases struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Items []Item `gorm:"foreignKey:CaseID"`
}
