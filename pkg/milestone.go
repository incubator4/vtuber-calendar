package pkg

import (
	"time"
)

type Milestone struct {
	ID        uint  `json:"id" gorm:"primarykey"`
	VtuberID  int64 `json:"vid" gorm:"column:vid"`
	Date      time.Time
	Event     string
	IsDisplay bool
	IsDeleted bool
}

func (Milestone) TableName() string {
	return "milestone"
}
