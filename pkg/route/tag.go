package route

import (
	"github.com/gin-gonic/gin"
	"github.com/incubator4/vtuber-calendar/internal/dao"
	"net/http"
)

func registerTag(g *gin.RouterGroup) {
	g.GET("", ListTags)
	//g.GET("/:uid", GetVtuber)

}

func ListTags(c *gin.Context) {
	tags, err := dao.ListEventTags()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": tags,
		})
	}
}
