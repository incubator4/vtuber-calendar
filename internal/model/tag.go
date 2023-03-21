package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Tag struct {
	ID        int        `json:"id" gorm:"primaryKey;"`
	Name      string     `json:"name" gorm:"column:name"`
	Calendars []Calendar `json:"-" gorm:"many2many:calendar_tags;"`
}

func (Tag) TableName() string {
	return "tag"
}

type TagIDArray []int

func (tg *TagIDArray) Value() (driver.Value, error) {
	return json.Marshal(*tg)
}

func (tg *TagIDArray) Scan(src interface{}) error {

	switch src := src.(type) {
	case []byte:
		return tg.scanBytes(src)
	case string:
		return tg.scanBytes([]byte(src))
	case nil:
		*tg = nil
		return nil
	}

	return fmt.Errorf("pq: cannot convert %T to Int32Array", src)
}

func (tg *TagIDArray) scanBytes(src []byte) error {
	return json.Unmarshal(src, tg)
}
