package dao

import (
	"github.com/incubator4/vtuber-calendar/internal/model"
)

func GetCharacter(_c model.Vtuber) *model.Vtuber {
	var c = new(model.Vtuber)
	DB.Table("character").Where(&_c).First(c)
	return c
}

func ListCharacter() []model.Vtuber {
	var characters []model.Vtuber
	DB.Table("character").Order("id").Find(&characters)
	return characters
}
