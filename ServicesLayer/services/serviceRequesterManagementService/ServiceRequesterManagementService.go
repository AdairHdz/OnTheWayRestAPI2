package serviceRequesterManagementService

import (
	"net/http"
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)


type ServiceRequesterManagementService struct{}

func (ServiceRequesterManagementService) Register(serviceRequester businessEntities.ServiceRequester) error {
	registryError := serviceRequester.Register()
	return registryError
}

func (ServiceRequesterManagementService) Find() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		var serviceRequester businessEntities.ServiceRequester
		searchError := serviceRequester.Find(serviceRequesterID)

		if searchError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateUserDTOAsResponse(serviceRequester.User, serviceRequesterID)
		context.JSON(http.StatusOK, response)
	}
}

func (ServiceRequesterManagementService) Update(serviceRequester businessEntities.ServiceRequester) error {	
	updateError := serviceRequester.Update()
	return updateError
}
