package dao

import (
	"github.com/incubator4/vtuber-calendar/internal/model"
)

func GetCharacter(_c model.Character) *model.Character {
	var c = new(model.Character)
	DB.Table("character").Where(&_c).First(c)
	return c
}

func ListCharacter() []model.Character {
	var characters []model.Character
	DB.Table("character").Order("id").Find(&characters)
	return characters
}
