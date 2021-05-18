package serviceRequestManagementService

import (
	"net/http"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ServiceRequestManagementService struct{}

func (ServiceRequestManagementService) Register() gin.HandlerFunc {
	return func(context *gin.Context){
		receivedData := dataTransferObjects.ReceivedServiceRequestDTO{}
		context.BindJSON(&receivedData)

		serviceRequestEntity := mappers.CreateServiceRequestEntity(receivedData)
		databaseError := serviceRequestEntity.Register()

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateServiceRequestDTOAsResponse(serviceRequestEntity)
		context.JSON(http.StatusCreated, response)
	}
}

func (ServiceRequestManagementService) Find() gin.HandlerFunc {
	return func(context *gin.Context){
		serviceRequestID, parsingError := uuid.FromString(context.Param("serviceRequestId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		var serviceRequest businessEntities.ServiceRequest
		databaseError := serviceRequest.Find(serviceRequestID)

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateServiceRequestDTOWithDetailsAsResponse(serviceRequest)
		context.JSON(http.StatusOK, response)


	}
}

func (ServiceRequestManagementService) Update() gin.HandlerFunc {
	return func(context *gin.Context){
		serviceRequestID, parsingError := uuid.FromString(context.Param("serviceRequestId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		var serviceRequest businessEntities.ServiceRequest
		databaseError := serviceRequest.Find(serviceRequestID)

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		serviceStatus := struct {
			ServiceStatus uint8 `json:"serviceStatus"`
		}{}
		
		context.BindJSON(&serviceStatus)

		serviceRequest.ServiceStatus = serviceStatus.ServiceStatus
		databaseError = serviceRequest.Update()

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		context.Status(http.StatusOK)
	}
}
