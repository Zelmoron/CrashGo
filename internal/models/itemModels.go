package models

type Item struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:255"`
	Cost   int    `gorm:"type:integer"`
	Image  string `gorm:"type:varchar(50)"`
	Type   string `gorm:"type:varchar(50)"`
	CaseID uint
}
