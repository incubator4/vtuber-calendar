package route

import (
	"fmt"
	ics "github.com/arran4/golang-ical"
	"github.com/gin-gonic/gin"
	"github.com/incubator4/vtuber-calendar/internal/dao"
	"net/http"
)

func registerICS(g *gin.RouterGroup) {
	//g.GET("", ListCalendars)
	g.GET("/:uid", GetICS)
}

func GetICS(c *gin.Context) {
	UID := c.Param("uid")
	calendars, _ := dao.ListCombineCalendars(
		dao.WithUID([]string{UID}),
		dao.WithOrder("id"),
	)

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	cal.SetTzid("Asia/Shanghai")

	for _, calendar := range calendars {
		e := ics.NewEvent(fmt.Sprintf("%d - %s", calendar.VtuberID, calendar.StartTime))
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
