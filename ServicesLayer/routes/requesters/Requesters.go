package requesters

import (	
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/controllers/addressController"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/controllers/serviceRequesterController"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/addressManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceRequesterManagementService"
	"github.com/gin-gonic/gin"
)

var (
	serviceRequesterMgtService = serviceRequesterManagementService.ServiceRequesterManagementService{}
	addressMgtService = addressManagementService.AddressManagementService{}	
)

func Routes(route *gin.RouterGroup) {
	requesters := route.Group("/requesters")
	{
		requesters.POST("/", serviceRequesterController.RegisterServiceRequester())	
		requesters.GET("/:requesterId", serviceRequesterController.FindServiceRequester())
		requesters.PATCH("/:requesterId", serviceRequesterController.UpdateServiceRequester()) //TODO: Fix
		requesters.POST("/:requesterId/addresses", addressController.RegisterAddress())
		requesters.GET("/:requesterId/addresses", addressController.FindAllAddressesOfServiceRequester())		
		requesters.GET("/:requesterId/requests", func(context *gin.Context) {
			// date := time.Date(2021, 5, 16)
			// repository := repositories.Repository{}
			//SELECT * FROM gorm.users WHERE created_at BETWEEN '2021-05-15 00:00:00' AND '2021-05-15 23:59:59';

		})
	}
}