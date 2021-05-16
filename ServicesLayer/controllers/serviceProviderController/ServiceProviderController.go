package serviceProviderController

import (
	"net/http"
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceProviderManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var (
	serviceProviderMgtService = serviceProviderManagementService.ServiceProviderManagementService{}
)


func RegisterServiceProvider() gin.HandlerFunc{
	return func(context *gin.Context){
		receivedData := dataTransferObjects.ReceivedUserDTO{}
		context.BindJSON(&receivedData)

		validate := validators.GetValidator()
		validationErrors := validate.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userEntity, mappingError := mappers.CreateUserEntity(receivedData, 1)

		if mappingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceProviderEntity := businessEntities.ServiceProvider{
			ID: uuid.NewV4(),
			User: userEntity,
			AverageScore: 0,
			PriceRates: nil,
		}

		serviceProviderMgtService := serviceProviderManagementService.ServiceProviderManagementService{}
		registryError := serviceProviderMgtService.Register(serviceProviderEntity)

		if registryError != nil {
			context.AbortWithStatus(http.StatusConflict)
		}

		response := mappers.CreateUserDTOAsResponse(serviceProviderEntity.User, serviceProviderEntity.ID)
		
		context.JSON(http.StatusCreated, response)
	}
}

func FindMatches() gin.HandlerFunc{
	return func(context *gin.Context){		
		// maxPriceRate, parsingError := strconv.ParseFloat(context.Param("maxPriceRate"), 32)

		// if parsingError != nil {
		// 	return
		// }

		// city := context.Param("city")
		// kindOfService, parsingError := strconv.ParseUint(context.Param("kindOfService"), 10, 8)

		// if parsingError != nil {
		// 	return
		// }

		// context.Params

	}
}

func Find() gin.HandlerFunc{
	return func(context *gin.Context){
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceProvider, searchError := serviceProviderMgtService.Find(serviceProviderID)

		if searchError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateUserDTOAsResponse(serviceProvider.User, serviceProvider.ID)
		context.JSON(http.StatusOK, response)
	}
}

func Update() gin.HandlerFunc{
	return func(context *gin.Context){
		
	}
}