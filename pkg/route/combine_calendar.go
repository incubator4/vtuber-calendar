package route

import (
	"github.com/gin-gonic/gin"
	"github.com/incubator4/vtuber-calendar/internal/dao"
	"net/http"
	"strconv"
)

func registerCombineCalendars(g *gin.RouterGroup) {
	g.GET("", ListCombineCalendars)
	g.GET("/:id", ValidateCalendarID, GetCombineCalendars)
}

func ListCombineCalendars(c *gin.Context) {
	timeRange, err := getStartAndEndOfDate(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	vids := c.QueryArray("vid")
	uids := c.QueryArray("uid")
	all, err := strconv.ParseBool(c.DefaultQuery("all", "false"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	calendars, err := dao.ListCombineCalendars(
		dao.WithUID(uids),
		dao.WithVID(vids),
		dao.WithTimeRange(timeRange),
		dao.WithOrder("id"),
		dao.WithAll(all),
		dao.Where("is_delete", false),
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": calendars,
	})

}

func GetCombineCalendars(c *gin.Context) {
	id := c.MustGet("id").(int)
	calendar := dao.GetCombineCalendar(dao.Preload("Tags"), dao.Where("id = ?", id))
	c.JSON(http.StatusOK, calendar)
}
