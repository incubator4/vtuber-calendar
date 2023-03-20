package dao

import "github.com/incubator4/vtuber-calendar/pkg"

func GetCharacter(_c pkg.Character) *pkg.Character {
	var c = new(pkg.Character)
	DB.Table("character").Where(&_c).First(c)
	return c
}

func ListCharacter() []pkg.Character {
	var characters []pkg.Character
	DB.Table("character").Order("id").Find(&characters)
	return characters
}
