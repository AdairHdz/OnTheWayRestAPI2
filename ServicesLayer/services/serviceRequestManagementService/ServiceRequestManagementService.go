package serviceRequestManagementService

import (
	"net/http"
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/customErrors"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ServiceRequestManagementService struct{}

func (ServiceRequestManagementService) Register() gin.HandlerFunc {
	return func(context *gin.Context){
		receivedData := dataTransferObjects.ReceivedServiceRequestDTO{}
		context.BindJSON(&receivedData)

		validator := validators.GetValidator()
		validationErrors := validator.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
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

func (ServiceRequestManagementService) FindByID() gin.HandlerFunc {
	return func(context *gin.Context){
		serviceRequestID, parsingError := uuid.FromString(context.Param("serviceRequestId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		var serviceRequest businessEntities.ServiceRequest
		databaseError := serviceRequest.Find(serviceRequestID)

		if databaseError != nil {
			_, errorIsOfTypeRecordNotFound := databaseError.(customErrors.RecordNotFoundError)
			if errorIsOfTypeRecordNotFound {
				context.AbortWithStatus(http.StatusNotFound)
				return
			}
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateServiceRequestDTOWithDetailsAsResponse(serviceRequest)
		context.JSON(http.StatusOK, response)


	}
}

func (ServiceRequestManagementService) FindByDate(userType int) gin.HandlerFunc {
	return func(context *gin.Context){
		dateOfServiceRequestsToBeFetched := context.Query("date")		

		_, dateParsingError := time.Parse("2006-01-02", dateOfServiceRequestsToBeFetched)
		if dateParsingError != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		
		
		var serviceRequests[] businessEntities.ServiceRequest
		
		var id uuid.UUID
		var parsingErrorUUID error		
		if userType == businessEntities.ServiceProviderType {						
			id, parsingErrorUUID = uuid.FromString(context.Param("providerId"))			
		}else {			
			id, parsingErrorUUID = uuid.FromString(context.Param("requesterId"))			
		}
		
		if parsingErrorUUID != nil {			
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		serviceRequestEntity := businessEntities.ServiceRequest{}
		serviceRequests, databaseError := serviceRequestEntity.FindByDate(dateOfServiceRequestsToBeFetched, id, userType)		
		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		var serviceRequestDTOs []dataTransferObjects.ResponseServiceRequestDTOWithDetails
		for _, serviceRequestElement := range serviceRequests {
			serviceRequestDTOs = append(serviceRequestDTOs, mappers.CreateServiceRequestDTOWithDetailsAsResponse(serviceRequestElement))
		}
		
		if len(serviceRequestDTOs) == 0 {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.JSON(http.StatusOK, serviceRequestDTOs)


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
