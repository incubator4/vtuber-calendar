package dao

import (
	"github.com/incubator4/vtuber-calendar/internal/model"
)

func ListCalendars(options ...Option) ([]model.CharacterCalendar, error) {
	var calendars []model.CharacterCalendar
	//db := params.TimeRange.DB(DB)
	db := DB
	for _, option := range options {
		db = option(db)
	}
	result := db.Find(&calendars)

	return calendars, result.Error
}

func GetCalendar(id int) *model.CharacterCalendar {
	var c = new(model.CharacterCalendar)
	DB.First(&c, id)
	return c
}

func UpdateCalendar(cal model.Calendar) *model.CharacterCalendar {
	DB.Save(&cal)
	var c = new(model.CharacterCalendar)
	DB.Where("id = ?", cal.ID).First(&c)
	return c
}

func CreateCalendar(cal model.Calendar) (*model.CharacterCalendar, error) {
	c := cal
	var err error
	err = DB.Create(&c).Error
	if err != nil {
		return nil, err
	}
	var cc = new(model.CharacterCalendar)
	err = DB.Where("id = ?", c.ID).First(&cc).Error
	return cc, err

}

func DeleteCalendar(id int) error {
	return DB.Table("calendar").Where("id = ?", id).Update("is_delete", true).Error
}
