package model

import (
	"time"
)

type Calendar struct {
	ID          int       `json:"id" gorm:"primaryKey;"`
	CharacterID int       `json:"cid" gorm:"column:cid;<-"`
	Title       string    `json:"title" gorm:"<-"`
	StartTime   time.Time `json:"start_time" gorm:"column:start_time;type:time;<-"`
	EndTime     time.Time `json:"end_time" gorm:"column:end_time;type:time;<-"`
	TagID       int       `json:"tag_id" gorm:"column:tag_id;<-"`
	IsActive    bool      `json:"is_active" gorm:"column:is_active;<-"`
}

type Character struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UID       int    `json:"uid" gorm:"column:uid"`
	LiveID    int    `json:"live_id" gorm:"column:live_id"`
	MainColor string `json:"main_color" gorm:"column:main_color"`
}

type CombineCalendar struct {
	Character
	Calendar
	Tags TagIDArray `json:"tags"`
}

//func (CombineCalendar) TableName() string {
//	return "calendar"
//}

func (Calendar) TableName() string {
	return "calendar"
}
