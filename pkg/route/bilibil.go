package route

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/biligo"
	"net/http"
	"strconv"
)

var bClient *biligo.CommClient

func init() {
	bClient = biligo.NewCommClient(&biligo.CommSetting{
		Client:    nil,
		DebugMode: false,
		UserAgent: "",
	})
}

func registerBilibili(g *gin.RouterGroup) {
	//g.GET("", ListVtubers)
	g.GET("/:uid", GetBiliUserInfo)

}

func GetBiliUserInfo(c *gin.Context) {

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	info, err := bClient.UserGetInfo(int64(uid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": info,
		})
	}

}
