package mappings
import (
	"gitlab.com/satvikshrivas26/olympic-fanzone/controllers"
	"github.com/gin-gonic/gin"
)
var Router *gin.Engine
func CreateUrlMappings() {
	Router = gin.Default()
	Router.Use(controllers.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		// v1.GET("/users/:id", controllers.GetUserDetail)
		v1.GET("/events", controllers.ScheduleAllEvents)
		v1.GET("/filter", controllers.FilteredSchedule)
		v1.GET("/live", controllers.LiveSchedule)
	
	}
}