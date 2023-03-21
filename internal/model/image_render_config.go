package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ImageConfig struct {
	URL    string `json:"url" gorm:"column:url"`
	Width  int    `json:"width" gorm:"column:width"`
	Height int    `json:"height" gorm:"column:height"`
}

type OffSet struct {
	X int `json:"x" gorm:"column:offset_x"`
	Y int `json:"y" gorm:"column:offset_y"`
}

type OffSetConfig struct {
	OffSet `gorm:"embedded"`
	Prefix string `json:"prefix" gorm:"column:prefix"`
	Suffix string `json:"suffix" gorm:"column:suffix"`
}

type FontConfig struct {
	Family string `json:"family" gorm:"column:family"`
	Size   int    `json:"size" gorm:"column:size"`
	Color  string `json:"color" gorm:"column:color"`
	Stroke string `json:"stroke" gorm:"column:stroke"`
	Layout string `json:"layout" gorm:"column:layout"`
	Style  string `json:"style" gorm:"column:style"`
}

type TextGroupArray [][]int

type TextGroupConfig struct {
	OffSet `json:"offset" gorm:"embedded;embeddedPrefix:group_"`
}

type ImageRenderConfig struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	Name            string `json:"name" gorm:"column:name"`
	ImageConfig     `json:"image" gorm:"embedded;embeddedPrefix:image_"`
	TextGroup       TextGroupArray `json:"text_group" gorm:"column:text_group;type:jsonb"`
	TextOffSet      OffSet         `json:"text_offset" gorm:"embedded;embeddedPrefix:text_"`
	TextGroupOffSet OffSet         `json:"text_group_offset" gorm:"embedded;embeddedPrefix:text_group_"`
	Row             OffSetConfig   `json:"row" gorm:"embedded;embeddedPrefix:row_"`
	Col             OffSetConfig   `json:"col" gorm:"embedded;embeddedPrefix:col_"`
	Font            FontConfig     `json:"font" gorm:"embedded;embeddedPrefix:font_"`
}

func (ImageRenderConfig) TableName() string {
	return "image_render_config"
}

func (tg *TextGroupArray) Value() (driver.Value, error) {
	return json.Marshal(*tg)
}

func (tg *TextGroupArray) Scan(src interface{}) error {

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

func (tg *TextGroupArray) scanBytes(src []byte) error {
	return json.Unmarshal(src, tg)
}
