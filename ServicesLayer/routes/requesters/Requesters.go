package requesters

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/middlewares"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/addressManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceRequestManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceRequesterManagementService"
	"github.com/gin-gonic/gin"
)

var (
	serviceRequesterMgtService = serviceRequesterManagementService.ServiceRequesterManagementService{}
	addressMgtService          = addressManagementService.AddressManagementService{}
	serviceRequestMgtService   = serviceRequestManagementService.ServiceRequestManagementService{}
)

func Routes(route *gin.RouterGroup) {
	requesters := route.Group("/requesters")
	{
		requesters.Use(middlewares.Authenticate())
		requesters.GET("/:requesterId", serviceRequesterMgtService.Find())
		requesters.PATCH("/:requesterId", serviceRequesterMgtService.Update())
		requesters.POST("/:requesterId/addresses", addressMgtService.Register())
		requesters.GET("/:requesterId/addresses", addressMgtService.FindAll())
		requesters.GET("/:requesterId/requests", serviceRequestMgtService.FindByDate(businessEntities.ServiceRequesterType))
		requesters.GET(":requesterId/statistics", serviceRequesterMgtService.GetStatistics())
	}
}
