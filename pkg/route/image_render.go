package route

import (
	"github.com/gin-gonic/gin"
	dao2 "github.com/incubator4/vtuber-calendar/internal/dao"
	"net/http"
	"strconv"
)

func registerImageRender(g *gin.RouterGroup) {
	g.GET("", ListImageRenderConfig)
	g.POST("", CreateCalendar)
	g.GET("/:id", ValidateCalendarID, GetImageRenderConfig)
	g.PUT("/:id", ValidateCalendarID, UpdateCalendar)
	g.DELETE("/:id", ValidateCalendarID, DeleteCalendar)
}

func ListImageRenderConfig(c *gin.Context) {
	stringOwner := c.Query("owner")

	owner, err := strconv.Atoi(stringOwner)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	configs, err := dao2.ListImageRenderConfig(
		dao2.WithOwner(owner),
		dao2.WithOrder("id"),
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": configs,
	})

}

func GetImageRenderConfig(c *gin.Context) {
	id := c.MustGet("id").(int)
	config := dao2.GetImageRenderConfig(id)
	c.JSON(http.StatusOK, config)
}
