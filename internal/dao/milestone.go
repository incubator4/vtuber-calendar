package dao

import (
	"github.com/incubator4/vtuber-calendar/pkg"
	"gorm.io/gorm"
)

func ListMileStone(options ...Option) ([]pkg.Milestone, error) {
	var milestones []pkg.Milestone
	//db := params.TimeRange.DB(DB)
	db := DB
	for _, option := range options {
		db = option(db)
	}
	result := db.Find(&milestones)

	return milestones, result.Error
}

func WithVtuberID(vid int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("vid = ?", vid)
	}
}
