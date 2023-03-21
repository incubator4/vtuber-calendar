package route

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	r.Use(HttpRequestTotalMetrics, HttpRequestDurationMetrics)
	api := r.Group("/api")

	routes := []struct {
		Group    *gin.RouterGroup
		Register func(g *gin.RouterGroup)
	}{
		{api.Group("/cal"), registerCalendars},
		{api.Group("/calendar"), registerCombineCalendars},
		{api.Group("/vtubers"), registerVtubers},
		{api.Group("/characters"), registerVtubers},
		{api.Group("/tags"), registerTag},
		{api.Group("/image_render"), registerImageRender},
		{api.Group("/ics"), registerICS},
		{api.Group("/status"), registerStatus},
		{api.Group("/milestones"), registerMilestone},
		{api.Group("/bilibili"), registerBilibili},
		{api.Group("/metrics"), registerMetrics},
	}

	for _, route := range routes {
		route.Register(route.Group)
	}

	return r
}
