package dao

import (
	"github.com/incubator4/vtuber-calendar/internal/model"
)

func ListCalendars(options ...Option) ([]model.CombineCalendar, error) {
	var calendars []model.CombineCalendar
	db := DB
	for _, option := range options {
		db = option(db)
	}

	result := db.Find(&calendars)

	return calendars, result.Error
}

func GetCalendar(options ...Option) *model.CombineCalendar {
	var c = new(model.CombineCalendar)
	db := DB
	for _, option := range options {
		db = option(db)
	}
	db.First(&c)
	return c
}

func UpdateCalendar(cal model.Calendar) *model.CombineCalendar {
	DB.Save(&cal)
	var c = new(model.CombineCalendar)
	DB.Where("id = ?", cal.ID).First(&c)
	return c
}

func CreateCalendar(cal model.Calendar) (*model.CombineCalendar, error) {
	c := cal
	var err error
	err = DB.Create(&c).Error
	if err != nil {
		return nil, err
	}
	var cc = new(model.CombineCalendar)
	err = DB.Where("id = ?", c.ID).First(&cc).Error
	return cc, err

}

func DeleteCalendar(id int) error {
	return DB.Table("calendar").Where("id = ?", id).Update("is_delete", true).Error
}
