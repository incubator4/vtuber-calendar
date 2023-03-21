package dao

import (
	"github.com/incubator4/vtuber-calendar/internal/model"
)

func ListTags(options ...Option) ([]model.Tag, error) {
	var tags []model.Tag
	db := DB
	for _, option := range options {
		db = option(db)
	}
	result := db.Debug().Find(&tags)

	return tags, result.Error
}

func GetTags(options ...Option) (*model.Tag, error) {
	var tags *model.Tag
	db := DB
	for _, option := range options {
		db = option(db)
	}
	result := db.Find(&tags)

	return tags, result.Error
}
