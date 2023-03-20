package dao

import "github.com/incubator4/vtuber-calendar/pkg"

func ListEventTags(options ...Option) ([]pkg.EventTag, error) {
	var tags []pkg.EventTag
	//db := params.TimeRange.DB(DB)
	db := DB
	for _, option := range options {
		db = option(db)
	}
	result := db.Find(&tags)

	return tags, result.Error
}
