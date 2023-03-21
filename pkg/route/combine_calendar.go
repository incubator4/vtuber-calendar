package route

import (
	"github.com/gin-gonic/gin"
	"github.com/incubator4/vtuber-calendar/internal/dao"
	"net/http"
)

func registerCombineCalendars(g *gin.RouterGroup) {
	//g.GET("", ListCombineCalendars)
	g.GET("/:id", ValidateCalendarID, GetCombineCalendars)
}

func GetCombineCalendars(c *gin.Context) {
	id := c.MustGet("id").(int)
	calendar := dao.GetCombineCalendar(dao.Preload("Tags"), dao.Where("id = ?", id))
	c.JSON(http.StatusOK, calendar)
}
