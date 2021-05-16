package providers

import (					
	"github.com/gin-gonic/gin"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceProviderManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/reviewManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/priceRateManagementService"
)

var(
	serviceProviderMgtService = serviceProviderManagementService.ServiceProviderManagementService{}
	reviewMgtService = reviewManagementService.ReviewManagementService{}
	priceRateMgtService = priceRateManagementService.PriceRateManagementService{}
)

func Routes(route *gin.RouterGroup) {
	providers := route.Group("/providers")
	{
		providers.POST("/", serviceProviderMgtService.Register())
		providers.GET("/", serviceProviderMgtService.FindMatches()) //TODO
		providers.GET("/:providerId", serviceProviderMgtService.Find())
		providers.PATCH("/:providerId", serviceProviderMgtService.Update())
		providers.PUT("/:providerId/image", func(context *gin.Context) {

		})

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