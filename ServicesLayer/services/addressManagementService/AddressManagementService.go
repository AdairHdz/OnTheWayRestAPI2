package addressManagementService

import (
	"net/http"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type AddressManagementService struct{}

func (AddressManagementService) Register() gin.HandlerFunc {
	return func(context *gin.Context) {

		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format.")
			return
		}

		receivedData := dataTransferObjects.ReceivedAddressDTO{}
		context.BindJSON(&receivedData)

		validator := validators.GetValidator()
		validationErrors := validator.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
			return
		}

		addressEntity := mappers.CreateAddressEntity(receivedData, serviceRequesterID)

		databaseError := addressEntity.Register()

		if databaseError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to register the resource.")
			return
		}

		response := mappers.CreateAddressDTOAsResponse(addressEntity)

		context.JSON(http.StatusCreated, response)
	}
}

func (AddressManagementService) FindAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format.")
			return
		}

		var address businessEntities.Address
		addresses, databaseError := address.FindAllAddressesOfServiceRequester(serviceRequesterID)

		if databaseError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to retrieve the data.")
			return
		}

		response := []dataTransferObjects.ResponseAddressWithCityDTO{}

		for _, address := range addresses {
			response = append(response, mappers.CreateAddressDTOWithCityAsResponse(address))
		}

		if len(addresses) == 0 {
			context.AbortWithStatusJSON(http.StatusNotFound, "There is no data for the period you provided.")
			return
		}
		context.JSON(http.StatusOK, response)
	}

}
