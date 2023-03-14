package route

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		cal := api.Group("/cal")
		registerCalendars(cal)
	}

	{
		vtuber := api.Group("/vtubers")
		characters := api.Group("/characters")
		registerVtubers(vtuber)
		registerVtubers(characters)
	}

	{
		eventTags := api.Group("/event_tags")
		registerEventTag(eventTags)
	}

	{
		imageRender := api.Group("/image_render")
		registerImageRender(imageRender)
	}

	{
		ics := api.Group("/ics")
		registerICS(ics)
	}

	{
		status := api.Group("/status")
		registerStatus(status)
	}

	{
		milestone := api.Group("/milestones")
		registerMilestone(milestone)
	}

	return r
}
