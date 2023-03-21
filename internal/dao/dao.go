package dao

import (
	"github.com/incubator4/vtuber-calendar/pkg/config"
	"github.com/incubator4/vtuber-calendar/pkg/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := config.GlobalConfig.DbUrl
	DB, err = gorm.Open(postgres.New(postgres.Config{

		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		panic(err)
	}

}

type ListCalendarParams struct {
	CIDArray  []string
	UIDArray  []string
	TimeRange types.TimeRange
}

type Option func(*gorm.DB) *gorm.DB

func WithTimeRange(tr types.TimeRange) Option {
	return func(db *gorm.DB) *gorm.DB {
		if tr.Start.IsZero() && tr.End.IsZero() {
			return db
		} else if tr.Start.IsZero() {
			return db.Where("start_time <= ?", tr.End)
		} else if tr.End.IsZero() {
			return db.Where("start_time >= ?", tr.Start)
		} else {
			return db.Where("start_time >= ? AND start_time <= ?", tr.Start, tr.End)
		}
	}
}

func WithUID(UIDArray []string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if UIDArray != nil && len(UIDArray) > 0 {
			return db.Where("uid IN (?)", UIDArray)
		} else {
			return db
		}
	}
}

func WithCID(CIDArray []string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if CIDArray != nil && len(CIDArray) > 0 {
			return db.Where("id IN (?)", CIDArray)
		} else {
			return db
		}
	}
}

func WithOrder(order string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func WithAll(all bool) Option {
	if all {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true)
	}

}

func Preload(query string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(query)
	}
}

func Where(query interface{}, args ...interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(query, args)
	}
}

func CombineCalendar(options ...Option) Option {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Table("vtuber_calendar")
		for _, option := range options {
			db = option(db)
		}
		return db
	}
}
