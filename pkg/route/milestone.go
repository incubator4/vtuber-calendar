package route

import (
	"github.com/gin-gonic/gin"
	"github.com/incubator4/vtuber-calendar/pkg/dao"
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

	milestones, err := dao.ListMileStone(dao.WithVtuberID(vid), dao.WithOrder("date"))
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
