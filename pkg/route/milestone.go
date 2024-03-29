package route

import (
	"github.com/gin-gonic/gin"
	dao2 "github.com/incubator4/vtuber-calendar/internal/dao"
	"net/http"
	"strconv"
)

func registerMilestone(g *gin.RouterGroup) {
	g.GET("/:vid", ListMilestone)

}

func ListMilestone(c *gin.Context) {
	stringVid := c.Param("vid")
	vid, err := strconv.Atoi(stringVid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}

	milestones, err := dao2.ListMileStone(dao2.WithVtuberID(vid), dao2.WithOrder("date"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": milestones,
	})
}
