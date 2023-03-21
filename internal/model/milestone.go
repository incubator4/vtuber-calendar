package model

import (
	"time"
)

type Milestone struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	VtuberID    int64     `json:"vid" gorm:"column:vid"`
	Date        time.Time `json:"date"`
	Event       string    `json:"event"`
	Description string    `json:"description"`
	IsDisplay   bool      `json:"is_display"`
	IsDeleted   bool      `json:"is_deleted"`
}

func (Milestone) TableName() string {
	return "milestone"
}
