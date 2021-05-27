package serviceRequests

import (
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/middlewares"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceRequestManagementService"
	"github.com/gin-gonic/gin"
)

var(
	serviceRequestMgtService = serviceRequestManagementService.ServiceRequestManagementService{}
)

func Routes(route *gin.RouterGroup) {
	serviceRequest := route.Group("/requests")
	{
		serviceRequest.Use(middlewares.Authenticate())
		serviceRequest.POST("/", serviceRequestMgtService.Register())
		serviceRequest.GET("/:serviceRequestId", serviceRequestMgtService.FindByID())
		serviceRequest.PATCH("/:serviceRequestId", serviceRequestMgtService.Update())
	}
}