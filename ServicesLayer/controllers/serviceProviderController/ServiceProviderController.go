package serviceProviderController

import (
	"net/http"	
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceProviderManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var (
	serviceProviderMgtService = serviceProviderManagementService.ServiceProviderManagementService{}
)


func RegisterServiceProvider() gin.HandlerFunc{
	return func(context *gin.Context){
		receivedData := struct {
			Names string `json:"names" validate:"required,min=1,max=50,lettersAndSpaces"`
			LastName string `json:"lastName" validate:"required,min=1,max=50,lettersAndSpaces"`
			EmailAddress string `json:"emailAddress" validate:"required,email,max=254"`
			Password string `json:"password" validate:"required,max=80"`
			StateID uuid.UUID `json:"stateId" validate:"required"`		
		}{}

		context.BindJSON(&receivedData)		
		validate := validators.GetValidator()
		validationErrors := validate.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		hashedPassword, hashingError := hashing.GenerateHash(receivedData.Password)

		if hashingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceProviderEntity := businessEntities.ServiceProvider{
			ID: uuid.NewV4(),
			User: businessEntities.User{
				ID: uuid.NewV4(),
				Names: receivedData.Names,
				LastName: receivedData.LastName,
				EmailAddress: receivedData.EmailAddress,
				Password: hashedPassword,
				UserType: 1,
				Verified: false,
				StateID: receivedData.StateID,				
			},
			AverageScore: 0,
			PriceRates: nil,
		}

		serviceProviderMgtService := serviceProviderManagementService.ServiceProviderManagementService{}
		registryError := serviceProviderMgtService.Register(serviceProviderEntity)

		if registryError != nil {
			context.AbortWithStatus(http.StatusConflict)
		}

		response := struct {
			ID uuid.UUID
			Names string
			LastName string
			EmailAddress string
			UserType uint8
			Verified bool
			StateID uuid.UUID
		}{
			ID: serviceProviderEntity.ID,
			Names: serviceProviderEntity.User.Names,
			LastName: serviceProviderEntity.User.LastName,
			EmailAddress: serviceProviderEntity.User.EmailAddress,
			UserType: serviceProviderEntity.User.UserType,
			Verified: serviceProviderEntity.User.Verified,
			StateID: serviceProviderEntity.User.StateID,
		}
		
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

		response := struct {
			ID uuid.UUID `json:"id"`
			Names string `json:"names"`
			LastName string `json:"lastName"`
			EmailAddress string `json:"emailAddress"`
			UserType uint8 `json:"userType"`
			Verified bool `json:"verified"`
			StateID uuid.UUID `json:"stateId"`
		}{
			ID: serviceProvider.ID,
			Names: serviceProvider.User.Names,
			LastName: serviceProvider.User.LastName,
			EmailAddress: serviceProvider.User.EmailAddress,
			UserType: serviceProvider.User.UserType,
			Verified: serviceProvider.User.Verified,
			StateID: serviceProvider.User.StateID,
		}

		context.JSON(http.StatusOK, response)
	}
}

func Update() gin.HandlerFunc{
	return func(context *gin.Context){
		
	}
}