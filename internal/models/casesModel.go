package models

type CasesModel struct {
	ID    int         `gorm:"primaryKey"`
	Name  string      `gorm:"size:255"`
	Image string      `gorm:"size:255"`
	Items []ItemModel `gorm:"foreignKey:CaseID"`
}

func (CasesModel) TableName() string {
	return "cases"
}
