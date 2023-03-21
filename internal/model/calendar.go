package model

import (
	"time"
)

type Calendar struct {
	ID        int       `json:"id" gorm:"primaryKey;"`
	VtuberID  int       `json:"vid" gorm:"column:vid;<-"`
	Title     string    `json:"title" gorm:"<-"`
	StartTime time.Time `json:"start_time" gorm:"column:start_time;type:time;<-"`
	EndTime   time.Time `json:"end_time" gorm:"column:end_time;type:time;<-"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active;<-"`
	Tags      []Tag     `json:"tags" gorm:"many2many:calendar_tags"`
}

type CombineCalendar struct {
	Vtuber
	Calendar
	ID int `json:"id" gorm:"primaryKey;"`
}

func (CombineCalendar) TableName() string {
	return "vtuber_calendar"
}

func (Calendar) TableName() string {
	return "calendar"
}
