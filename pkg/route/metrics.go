package route

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func registerMetrics(g *gin.RouterGroup) {
	g.GET("", gin.WrapH(promhttp.Handler()))
}
