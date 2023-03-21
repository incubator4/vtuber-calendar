package model

import (
	"time"
)

type Calendar struct {
	ID        int       `json:"id" gorm:"primaryKey;"`
	VtuberID  int       `json:"vid" gorm:"index,column:vid;<-"`
	Title     string    `json:"title" gorm:"<-"`
	StartTime time.Time `json:"start_time" gorm:"column:start_time;type:time;<-"`
	EndTime   time.Time `json:"end_time" gorm:"column:end_time;type:time;<-"`
	TagID     int       `json:"tag_id" gorm:"column:tag_id;<-"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active;<-"`
}

type CombineCalendar struct {
	Vtuber
	Calendar
	Tags TagIDArray `json:"tags"`
}

func (CombineCalendar) TableName() string {
	return "calendar"
}

func (Calendar) TableName() string {
	return "calendar"
}
