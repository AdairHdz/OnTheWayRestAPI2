package states

import (

	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/stateManagementService"
	"github.com/gin-gonic/gin"
)

var (
	_stateManagementService = stateManagementService.StateManagementService{}
)

func Routes(route *gin.RouterGroup) {
	states := route.Group("/states")
	{
		states.GET("/", _stateManagementService.FindAll())

		states.GET("/:stateId/cities", _stateManagementService.FindAllCitiesOfState())
	}
}