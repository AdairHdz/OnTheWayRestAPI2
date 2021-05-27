package priceRateManagementService

import (
	"net/http"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/customErrors"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
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

		validator := validators.GetValidator()
		validationErrors := validator.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		priceRateEntity, mappingError := mappers.CreatePriceRateEntity(receivedData, serviceProviderID)

		if mappingError != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		databaseError := priceRateEntity.Register()

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreatePriceRateDTOAsResponse(priceRateEntity)		
		context.JSON(http.StatusCreated, response)
	}
}

func (PriceRateManagementService) FindAll() gin.HandlerFunc {
	return func(context *gin.Context){
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		priceRate := businessEntities.PriceRate{}
		priceRates, databaseError := priceRate.Find(serviceProviderID)
		
		if databaseError != nil {
			context.Status(http.StatusConflict)
			return
		}
		
		response := mappers.CreatePriceRateDTOSliceAsResponse(priceRates)
		if len(response) == 0 {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}
		context.JSON(http.StatusOK, response)
	}
}

func (PriceRateManagementService) Delete() gin.HandlerFunc {
	return func(context *gin.Context){		
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		priceRateID, parsingError := uuid.FromString(context.Param("priceRateId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		priceRate := businessEntities.PriceRate{
			ID: priceRateID,
		}

		databaseError := priceRate.Delete(serviceProviderID)

		if databaseError != nil {			
			_, errorIsOfTypeRecordNotFound := databaseError.(customErrors.RecordNotFoundError)
			if errorIsOfTypeRecordNotFound {
				context.Status(http.StatusNotFound)
				return
			}
			context.Status(http.StatusConflict)
			return
		}

		context.Status(http.StatusNoContent)


	}
}
