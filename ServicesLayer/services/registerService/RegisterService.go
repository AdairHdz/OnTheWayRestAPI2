package registerService

import (
	"net/http"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)


type RegisterService struct{}

func (RegisterService) RegisterUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		receivedData := dataTransferObjects.ReceivedUserDTO{}
			context.BindJSON(&receivedData)
	
			validate := validators.GetValidator()
			validationErrors := validate.Struct(receivedData)
	
			if validationErrors != nil {
				context.AbortWithStatus(http.StatusBadRequest)
				return
			}
	
			userEntity, mappingError := mappers.CreateUserEntity(receivedData)
	
			if mappingError != nil {
				context.AbortWithStatus(http.StatusConflict)
				return
			}

			var response dataTransferObjects.ResponseUserDTO

			if receivedData.UserType == businessEntities.ServiceProviderType {
				serviceProviderEntity := businessEntities.ServiceProvider{
					ID: uuid.NewV4(),
					User: userEntity,
					AverageScore: 0,
					PriceRates: nil,
				}

				registryError := serviceProviderEntity.Register()
								
				if registryError != nil {
					context.AbortWithStatus(http.StatusConflict)
					return
				}

				response = mappers.CreateUserDTOAsResponse(serviceProviderEntity.User, serviceProviderEntity.ID)
			}else{				
				serviceRequesterEntity := businessEntities.ServiceRequester{
					ID: uuid.NewV4(),
					User: userEntity,
					Addresses: nil,
				}

				registryError := serviceRequesterEntity.Register()
				
				if registryError != nil {					
					context.AbortWithStatus(http.StatusConflict)
					return
				}

				response = mappers.CreateUserDTOAsResponse(serviceRequesterEntity.User, serviceRequesterEntity.ID)
			}
							
			context.JSON(http.StatusCreated, response)
	}
}