package route

import (
	"github.com/gin-gonic/gin"
	"github.com/incubator4/vtuber-calendar/internal/dao"
	"net/http"
	"strconv"
)

func registerTag(g *gin.RouterGroup) {
	g.GET("", ListTags)
	g.GET("/:id", GetTag)
}

func ListTags(c *gin.Context) {
	tags, err := dao.ListTags(dao.Preload("Calendars"))
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

func GetTag(c *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	tag, err := dao.GetTags(dao.Where("id = ?", id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": tag,
		})
	}
}
