package dao

import (
	"github.com/incubator4/vtuber-calendar/internal/model"
)

func ListImageRenderConfig(options ...Option) ([]model.ImageRenderConfig, error) {
	var configs []model.ImageRenderConfig
	db := DB
	for _, option := range options {
		db = option(db)
	}
	result := db.Find(&configs)

	return configs, result.Error
}

func GetImageRenderConfig(id int) *model.ImageRenderConfig {
	var c = new(model.ImageRenderConfig)
	DB.First(&c, id)
	return c
}
