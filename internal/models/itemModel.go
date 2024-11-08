package models

type ItemModel struct {
	ID     int    `gorm:"primaryKey"`
	Name   string `gorm:"size:255"`
	Cost   int    `gorm:"type:integer"`
	Image  string `gorm:"type:varchar(50)"`
	Type   string `gorm:"type:varchar(50)"`
	CaseID int
}

func (ItemModel) TableName() string {
	return "items"
}
