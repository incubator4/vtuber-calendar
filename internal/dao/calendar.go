package dao

import "github.com/incubator4/vtuber-calendar/pkg"

func ListCalendars(options ...Option) ([]pkg.CharacterCalendar, error) {
	var calendars []pkg.CharacterCalendar
	//db := params.TimeRange.DB(DB)
	db := DB
	for _, option := range options {
		db = option(db)
	}
	result := db.Find(&calendars)

	return calendars, result.Error
}

func GetCalendar(id int) *pkg.CharacterCalendar {
	var c = new(pkg.CharacterCalendar)
	DB.First(&c, id)
	return c
}

func UpdateCalendar(cal pkg.Calendar) *pkg.CharacterCalendar {
	DB.Save(&cal)
	var c = new(pkg.CharacterCalendar)
	DB.Where("id = ?", cal.ID).First(&c)
	return c
}

func CreateCalendar(cal pkg.Calendar) (*pkg.CharacterCalendar, error) {
	c := cal
	var err error
	err = DB.Create(&c).Error
	if err != nil {
		return nil, err
	}
	var cc = new(pkg.CharacterCalendar)
	err = DB.Where("id = ?", c.ID).First(&cc).Error
	return cc, err

}

func DeleteCalendar(id int) error {
	return DB.Table("calendar").Where("id = ?", id).Update("is_delete", true).Error
}
