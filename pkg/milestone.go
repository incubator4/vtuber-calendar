package pkg

import (
	"time"
)

type Milestone struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	VtuberID  int64     `json:"vid" gorm:"column:vid"`
	Date      time.Time `json:"date"`
	Event     string    `json:"event"`
	IsDisplay bool      `json:"is_display"`
	IsDeleted bool      `json:"is_deleted"`
}

func (Milestone) TableName() string {
	return "milestone"
}
