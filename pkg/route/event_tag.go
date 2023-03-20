package route

import (
	"github.com/gin-gonic/gin"
	"github.com/incubator4/vtuber-calendar/internal/dao"
	"net/http"
)

func registerEventTag(g *gin.RouterGroup) {
	g.GET("", ListEventTags)
	//g.GET("/:uid", GetVtuber)

}

func ListEventTags(c *gin.Context) {
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

//func GetVtuber(c *gin.Context) {
//	stringUID := c.Param("uid")
//	uid, err := strconv.Atoi(stringUID)
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"error": err,
//		})
//		return
//	}
//	character := dao.GetCharacter(pkg.Character{UID: uid})
//	c.JSON(http.StatusOK, gin.H{
//		"data": character,
//	})
//}
