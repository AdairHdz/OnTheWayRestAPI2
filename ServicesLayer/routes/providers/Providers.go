package providers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/controllers/serviceProviderController"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/priceRateManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/reviewManagementService"
	"github.com/gin-gonic/gin"
)

var(	
	reviewMgtService = reviewManagementService.ReviewManagementService{}
	priceRateMgtService = priceRateManagementService.PriceRateManagementService{}
)

func Routes(route *gin.RouterGroup) {
	providers := route.Group("/providers")
	{
		providers.POST("/", serviceProviderController.RegisterServiceProvider())
		providers.GET("/", serviceProviderController.FindMatches()) //TODO
		providers.GET("/:providerId", serviceProviderController.Find())
		providers.PATCH("/:providerId", serviceProviderController.Update())
		providers.PUT("/:providerId/image", serviceProviderController.UpdateServiceProviderImage())

		reviews := providers.Group("/:providerId")
		{
			reviews.POST("/reviews", reviewMgtService.Register())
			reviews.GET("/reviews", reviewMgtService.Find())
			reviews.POST("/reviews/:reviewId/evidence", func(context *gin.Context) {

			})
		}

		priceRates := providers.Group("/:providerId")
		{
			priceRates.POST("/priceRates", priceRateMgtService.Register())
			priceRates.GET("/priceRates", priceRateMgtService.FindAll())
			priceRates.DELETE("/priceRates/:priceRateId", priceRateMgtService.Delete())
		}
		
		requests := providers.Group("/:providerId")
		{
			requests.GET("/requests", func(context *gin.Context) {
				
			})
		}
	}
}