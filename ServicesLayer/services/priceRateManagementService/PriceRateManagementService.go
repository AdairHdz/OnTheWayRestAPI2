package priceRateManagementService

import (	
	"net/http"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)


type PriceRateManagementService struct{}

func (PriceRateManagementService) Register() gin.HandlerFunc {
	return func(context *gin.Context){
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		receivedData := dataTransferObjects.ReceivedPriceRateDTO{}

		context.BindJSON(&receivedData)

		priceRateEntity := mappers.CreatePriceRateEntity(receivedData, serviceProviderID)
		databaseError := priceRateEntity.Register()

		if databaseError != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		response := mappers.CreatePriceRateDTOAsResponse(priceRateEntity)	
		
		
		context.JSON(http.StatusCreated, response)
	}
}

func (PriceRateManagementService) FindAll() gin.HandlerFunc {
	return func(context *gin.Context){
		
	}
}

func (PriceRateManagementService) Delete() gin.HandlerFunc {
	return func(context *gin.Context){		
		
	}
}
