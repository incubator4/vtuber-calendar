package dao

import (
	"github.com/incubator4/vtuber-calendar/internal/model"
)

func GetCharacter(_c model.Vtuber) *model.Vtuber {
	var c = new(model.Vtuber)
	DB.Where(&_c).First(c)
	return c
}

func ListCharacter() []model.Vtuber {
	var characters []model.Vtuber
	DB.Order("id").Find(&characters)
	return characters
}

func ListVIDS(UIDArray []string) []string {
	var VIDArray []string
	DB.Table("vtuber").Select("id").Where("uid IN (?)", UIDArray).Find(&VIDArray)
	return VIDArray
}
