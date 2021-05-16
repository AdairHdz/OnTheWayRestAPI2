package addressController

import (
	"net/http"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/addressManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)


func RegisterAddress() gin.HandlerFunc{
	return func(context *gin.Context){

		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		receivedData := dataTransferObjects.ReceivedAddressDTO{}
		context.BindJSON(&receivedData)

		validator := validators.GetValidator()
		validationErrors := validator.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		addressEntity := mappers.CreateAddressEntity(receivedData, serviceRequesterID)

		addressMgtService := addressManagementService.AddressManagementService{}
		databaseError := addressMgtService.Register(addressEntity)

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateAddressDTOAsResponse(addressEntity)

		context.JSON(http.StatusCreated, response)
	}
}

func FindAllAddressesOfServiceRequester() gin.HandlerFunc {
	return func(context *gin.Context){
		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}
		
		addressMgtService := addressManagementService.AddressManagementService{}
		addresses, databaseError := addressMgtService.FindAll(serviceRequesterID)

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := []dataTransferObjects.ResponseAddressWithCityDTO{}

		for _, address := range addresses {
			response = append(response, mappers.CreateAddressDTOWithCityAsResponse(address))
		}
		context.JSON(http.StatusOK, response)
	}
}