package serviceRequesterController

import (
	"net/http"
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceRequesterManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var (
	serviceRequesterMgtService = serviceRequesterManagementService.ServiceRequesterManagementService{}
)

func RegisterServiceRequester() gin.HandlerFunc{
	return func(context *gin.Context) {
		receivedData := dataTransferObjects.ReceivedUserDTO{}
		context.BindJSON(&receivedData)

		validate := validators.GetValidator()
		validationErrors := validate.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		userEntity, mappingError := mappers.CreateUserEntity(receivedData, 2)

		if mappingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceRequesterEntity := businessEntities.ServiceRequester{
			ID: uuid.NewV4(),
			User: userEntity,
			Addresses: nil,
		}

		registryError := serviceRequesterMgtService.Register(serviceRequesterEntity)

		if registryError != nil {
			context.AbortWithStatus(http.StatusConflict)
		}

		response := mappers.CreateUserDTOAsResponse(serviceRequesterEntity.User, serviceRequesterEntity.ID)

		context.JSON(http.StatusCreated, response)
		
	}	
}

func FindServiceRequester() gin.HandlerFunc{
	return func(context *gin.Context){
		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceRequester, searchError := serviceRequesterMgtService.Find(serviceRequesterID)

		if searchError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateUserDTOAsResponse(serviceRequester.User, serviceRequesterID)
		context.JSON(http.StatusOK, response)

	}
}

func UpdateServiceRequester() gin.HandlerFunc{
	return func(context *gin.Context){
		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		receivedData := struct{
			Names string `json:"names"`
			LastName string `json:"lastName"`
			Password string `json:"password"`
		}{}

		bindingError := context.BindJSON(&receivedData)

		if bindingError != nil {
			return
		}

		serviceRequester := businessEntities.ServiceRequester{ }
		serviceRequester.Find(serviceRequesterID)

		if receivedData.Names != ""{
			serviceRequester.User.Names = receivedData.Names
		}

		if receivedData.LastName != "" {
			serviceRequester.User.LastName = receivedData.LastName
		}
				
		if receivedData.Password != "" {
			hashedPassword, hashingError := hashing.GenerateHash(serviceRequester.User.Password)	
			if hashingError != nil {
				context.AbortWithStatus(http.StatusConflict)
				return
			}
	
			serviceRequester.User.Password = hashedPassword
		}				
		
		validator :=  validators.GetValidator()
		validationErrors := validator.Var(serviceRequester.User.Names, "required,min=1,max=50,lettersAndSpaces")

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		validationErrors = validator.Var(serviceRequester.User.LastName, "required,min=1,max=50,lettersAndSpaces")

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		validationErrors = validator.Var(serviceRequester.User.Password, "required,max=80")

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		repository := repositories.Repository{}
		updateError := repository.Update(&serviceRequester.User)

		if updateError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		context.Status(http.StatusOK)

	}
}