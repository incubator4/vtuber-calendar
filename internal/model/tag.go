package model

type Tag struct {
	ID   int    `json:"id" gorm:"primaryKey;"`
	Name string `json:"name" gorm:"column:name"`
}

func (Tag) TableName() string {
	return "tag"
}
