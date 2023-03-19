package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mssola/useragent"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"strconv"
	"time"
)

var httpRequestsTotal *prometheus.CounterVec
var httpRequestDuration *prometheus.HistogramVec

var APP = "vtuber-calendar"

func init() {
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Vtuber API Total number of HTTP requests.",
	}, []string{"path", "method", "status", "app", "client_ip", "os", "model", "platform"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_requests_duration_seconds",
		Help:    "Duration of HTTP requests.",
		Buckets: []float64{0.1, 0.5, 1, 2.5, 5, 10},
	}, []string{"path", "method", "status", "app", "client_ip", "os", "model", "platform"})
}

type Middleware func(handler gin.HandlerFunc) gin.HandlerFunc

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	//c.Request.Cookies()
	if token == "example-token" {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
	}
}

func ValidateCalendarID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error": fmt.Sprintf("Invalid calendar ID %s", err),
		})

	} else {
		c.Set("id", id)
		c.Next()
	}
}

func HttpRequestTotalMetrics(c *gin.Context) {
	c.Next()
	uaString := c.GetHeader("User-Agent")
	ua := useragent.New(uaString)
	httpRequestsTotal.WithLabelValues(
		c.FullPath(),
		c.Request.Method,
		strconv.Itoa(c.Writer.Status()),
		APP,
		c.ClientIP(),
		ua.OS(),
		ua.Model(),
		ua.Platform(),
	).Inc()
}

func HttpRequestDurationMetrics(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	elapsedTime := time.Since(startTime).Seconds()
	uaString := c.GetHeader("User-Agent")
	ua := useragent.New(uaString)
	httpRequestDuration.WithLabelValues(
		c.FullPath(),
		c.Request.Method,
		strconv.Itoa(c.Writer.Status()),
		APP,
		c.ClientIP(),
		ua.OS(),
		ua.Model(),
		ua.Platform(),
	).Observe(elapsedTime)
}
