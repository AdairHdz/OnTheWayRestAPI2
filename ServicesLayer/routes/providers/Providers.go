package providers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/middlewares"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/priceRateManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/reviewManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceProviderManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceRequestManagementService"
	"github.com/gin-gonic/gin"
)

var (
	_reviewManagementService          = reviewManagementService.ReviewManagementService{}
	priceRateMgtService               = priceRateManagementService.PriceRateManagementService{}
	_serviceProviderManagementService = serviceProviderManagementService.ServiceProviderManagementService{}
	_serviceRequestManagementService  = serviceRequestManagementService.ServiceRequestManagementService{}
)

func Routes(route *gin.RouterGroup) {
	providers := route.Group("/providers")
	{
		providers.Use(middlewares.Authenticate())
		providers.GET("", _serviceProviderManagementService.FindMatches())
		providers.GET("/:providerId", _serviceProviderManagementService.Find())
		providers.PATCH("/:providerId", _serviceProviderManagementService.Update())
		providers.GET("/:providerId/statistics", _serviceProviderManagementService.GetStatistics())
		providers.PUT("/:providerId/image", _serviceProviderManagementService.UpdateServiceProviderImage())

		reviews := providers.Group("/:providerId")
		{
			reviews.POST("/reviews", _reviewManagementService.Register())
			reviews.GET("/reviews", _reviewManagementService.Find())
			reviews.POST("/reviews/:reviewId/evidence", _reviewManagementService.UploadEvidence())
		}

		priceRates := providers.Group("/:providerId")
		{
			priceRates.POST("/priceRates", priceRateMgtService.Register())
			priceRates.GET("/priceRates", priceRateMgtService.FindAll())
			priceRates.DELETE("/priceRates/:priceRateId", priceRateMgtService.Delete())
		}

		requests := providers.Group("/:providerId")
		{
			requests.GET("/requests", _serviceRequestManagementService.FindByDate(businessEntities.ServiceProviderType))
		}
	}
}
