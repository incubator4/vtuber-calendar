package route

import (
	"github.com/gin-gonic/gin"
	"github.com/incubator4/vtuber-calendar/internal/dao"
	"github.com/incubator4/vtuber-calendar/pkg"
	"net/http"
	"strconv"
)

func registerVtubers(g *gin.RouterGroup) {
	g.GET("", ListVtubers)
	g.GET("/:uid", GetVtuber)

}

func ListVtubers(c *gin.Context) {
	characters := dao.ListCharacter()
	c.JSON(http.StatusOK, gin.H{
		"data": characters,
	})
}

func GetVtuber(c *gin.Context) {
	stringUID := c.Param("uid")
	uid, err := strconv.Atoi(stringUID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	character := dao.GetCharacter(pkg.Character{UID: uid})
	c.JSON(http.StatusOK, gin.H{
		"data": character,
	})
}
