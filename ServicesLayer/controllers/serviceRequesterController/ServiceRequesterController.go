package serviceRequesterController

import (
	"net/http"
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceRequesterManagementService"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
)

var (
	serviceRequesterMgtService = serviceRequesterManagementService.ServiceRequesterManagementService{}
)

func RegisterServiceRequester() gin.HandlerFunc{
	return func(context *gin.Context) {
		receivedData := struct {
			Names string `json:"names" validate:"required,min=1,max=50,lettersAndSpaces"`
			LastName string `json:"lastName" validate:"required,min=1,max=50,lettersAndSpaces"`
			EmailAddress string `json:"emailAddress" validate:"required,email,max=254"`
			Password string `json:"password" validate:"required,max=80"`
			StateID uuid.UUID `json:"stateId" validate:"required"`		
		}{}

		context.BindJSON(&receivedData)
		var validate *validator.Validate = validator.New()
		validate.RegisterValidation("lettersAndSpaces", validators.LettersAndSpaces)
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

		serviceRequesterEntity := businessEntities.ServiceRequester{
			ID: uuid.NewV4(),
			User: businessEntities.User{
				ID: uuid.NewV4(),
				Names: receivedData.Names,
				LastName: receivedData.LastName,
				EmailAddress: receivedData.EmailAddress,
				Password: hashedPassword,
				UserType: 0,
				Verified: false,
				StateID: receivedData.StateID,				
			},
			Addresses: nil,
		}

		registryError := serviceRequesterMgtService.Register(serviceRequesterEntity)
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
			ID: serviceRequesterEntity.ID,
			Names: serviceRequesterEntity.User.Names,
			LastName: serviceRequesterEntity.User.LastName,
			EmailAddress: serviceRequesterEntity.User.EmailAddress,
			UserType: serviceRequesterEntity.User.UserType,
			Verified: serviceRequesterEntity.User.Verified,
			StateID: serviceRequesterEntity.User.StateID,
		}
		
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

		response := struct {
			ID uuid.UUID `json:"id"`
			Names string `json:"names"`
			LastName string `json:"lastName"`
			EmailAddress string `json:"emailAddress"`
			UserType uint8 `json:"userType"`
			Verified bool `json:"verified"`
			StateID uuid.UUID `json:"stateId"`
		}{
			ID: serviceRequester.ID,
			Names: serviceRequester.User.Names,
			LastName: serviceRequester.User.LastName,
			EmailAddress: serviceRequester.User.EmailAddress,
			UserType: serviceRequester.User.UserType,
			Verified: serviceRequester.User.Verified,
			StateID: serviceRequester.User.StateID,
		}

		context.JSON(http.StatusOK, response)

	}
}

// func UpdateServiceRequester() gin.HandlerFunc{
// 	return func(context *gin.Context){
// 		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

// 		if parsingError != nil {
// 			context.AbortWithStatus(http.StatusConflict)
// 			return
// 		}

// 		receivedData := struct{
// 			Names string `json:"names" validate:"required,min=1,max=50,lettersAndSpaces"`
// 			LastName string `json:"lastName" validate:"required,min=1,max=50,lettersAndSpaces"`
// 			Password string `json:"password" validate:"required,max=80"`
// 		}{}

// 		context.BindJSON(&receivedData)

// 		var validate *validator.Validate = validator.New()
// 		validate.RegisterValidation("lettersAndSpaces", validators.LettersAndSpaces)
// 		validationErrors := validate.Struct(receivedData)		

// 		if validationErrors != nil {
// 			context.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
		

// 		serviceRequester := businessEntities.ServiceRequester{ }
// 		serviceRequester.Find(serviceRequesterID)
// 		serviceRequester.User.Names = receivedData.Names
// 		serviceRequester.User.LastName = receivedData.LastName

// 		if receivedData.Password != "" {
// 			hashedPassword, hashingError := hashing.GenerateHash(serviceRequester.User.Password)	
// 			if hashingError != nil {
// 				context.AbortWithStatus(http.StatusConflict)
// 				return
// 			}
	
// 			serviceRequester.User.Password = hashedPassword
// 		}				
		
// 		repository := repositories.Repository{}
// 		updateError := repository.Update(&serviceRequester.User)

// 		if updateError != nil {
// 			context.AbortWithStatus(http.StatusConflict)
// 			return
// 		}

// 		context.Status(http.StatusOK)

// 	}
// }