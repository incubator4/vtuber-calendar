package model

type Vtuber struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UID       int    `json:"uid" gorm:"column:uid"`
	LiveID    int    `json:"live_id" gorm:"column:live_id"`
	MainColor string `json:"main_color" gorm:"column:main_color"`
}

func (Vtuber) TableName() string {
	return "vtuber"
}
