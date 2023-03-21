package route

import (
	"fmt"
	ics "github.com/arran4/golang-ical"
	"github.com/gin-gonic/gin"
	"github.com/incubator4/vtuber-calendar/internal/dao"
	"net/http"
	"strconv"
)

func registerICS(g *gin.RouterGroup) {
	//g.GET("", ListCalendars)
	g.GET("/:uid", GetICS)
}

func GetICS(c *gin.Context) {
	UID := c.Param("uid")
	all, err := strconv.ParseBool(c.DefaultQuery("all", "false"))
	if err != nil {
		all = false
	}
	calendars, _ := dao.ListCalendars(
		dao.CombineCalendar(dao.WithUID([]string{UID}), dao.WithOrder("id"), dao.WithAll(all)),
	)

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	cal.SetTzid("Asia/Shanghai")

	for _, calendar := range calendars {
		e := ics.NewEvent(fmt.Sprintf("%d - %s", calendar.CharacterID, calendar.StartTime))
		e.SetCreatedTime(calendar.StartTime)
		e.SetStartAt(calendar.StartTime)
		e.SetEndAt(calendar.EndTime)
		e.SetURL(fmt.Sprintf("bilibili://live/%d", calendar.LiveID))
		e.SetSummary(fmt.Sprintf("%s %s", calendar.Name, calendar.Title))
		e.SetDescription(fmt.Sprintf("直播间: https://live.bilibili.com/%d \n"+
			"个人空间: https://space.bilibili.com/%d", calendar.LiveID, calendar.UID))
		cal.AddVEvent(e)
	}
	c.String(http.StatusOK, cal.Serialize())
}
