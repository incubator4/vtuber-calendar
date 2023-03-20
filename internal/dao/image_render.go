package dao

import "github.com/incubator4/vtuber-calendar/pkg"

func ListImageRenderConfig(options ...Option) ([]pkg.ImageRenderConfig, error) {
	var configs []pkg.ImageRenderConfig
	db := DB
	for _, option := range options {
		db = option(db)
	}
	result := db.Find(&configs)

	return configs, result.Error
}

func GetImageRenderConfig(id int) *pkg.ImageRenderConfig {
	var c = new(pkg.ImageRenderConfig)
	DB.First(&c, id)
	return c
}
