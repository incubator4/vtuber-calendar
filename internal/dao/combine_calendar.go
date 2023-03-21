package dao

import "github.com/incubator4/vtuber-calendar/internal/model"

func ListCombineCalendars(options ...Option) ([]model.CombineCalendar, error) {
	var calendars []model.CombineCalendar
	db := DB
	for _, option := range options {
		db = option(db)
	}

	result := db.Find(&calendars)

	return calendars, result.Error
}

func GetCombineCalendar(options ...Option) *model.CombineCalendar {
	var c = new(model.CombineCalendar)
	db := DB
	for _, option := range options {
		db = option(db)
	}
	db.First(&c)
	return c
}
