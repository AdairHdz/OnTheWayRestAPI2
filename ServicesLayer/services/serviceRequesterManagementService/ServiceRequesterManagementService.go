package serviceRequesterManagementService

import (
	"net/http"
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/customErrors"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ServiceRequesterManagementService struct{}

func (ServiceRequesterManagementService) Find() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format.")
			return
		}

		var serviceRequester businessEntities.ServiceRequester
		searchError := serviceRequester.Find(serviceRequesterID)

		if searchError != nil {
			_, errorIsOfTypeRecordNotFoundError := searchError.(customErrors.RecordNotFoundError)
			if errorIsOfTypeRecordNotFoundError {
				context.AbortWithStatusJSON(http.StatusNotFound, "There are no matches for the ID you provided.")
				return
			}
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to retrieve the data.")
			return
		}

		response := mappers.CreateUserDTOAsResponse(serviceRequester.User, serviceRequesterID)
		context.JSON(http.StatusOK, response)
	}
}

func (ServiceRequesterManagementService) Update() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format.")
			return
		}

		receivedData := struct {
			Names    string `json:"names"`
			LastName string `json:"lastName"`
			Password string `json:"password"`
		}{}

		bindingError := context.BindJSON(&receivedData)

		if bindingError != nil {
			return
		}

		serviceRequester := businessEntities.ServiceRequester{}
		serviceRequester.Find(serviceRequesterID)

		if receivedData.Names != "" {
			serviceRequester.User.Names = receivedData.Names
		}

		if receivedData.LastName != "" {
			serviceRequester.User.LastName = receivedData.LastName
		}

		if receivedData.Password != "" {
			serviceRequester.User.Password = receivedData.Password
		}

		validator := validators.GetValidator()
		validationErrors := validator.Var(serviceRequester.User.Names, "required,min=1,max=50,lettersAndSpaces")

		if validationErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
			return
		}

		validationErrors = validator.Var(serviceRequester.User.LastName, "required,min=1,max=50,lettersAndSpaces")

		if validationErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
			return
		}

		validationErrors = validator.Var(serviceRequester.User.Password, "required,max=80")

		if validationErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
			return
		}

		hashedPassword, hashingError := hashing.GenerateHash(serviceRequester.User.Password)
		if hashingError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to update the resource.")
			return
		}
		serviceRequester.User.Password = hashedPassword

		updateError := serviceRequester.Update()

		if updateError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to update the resource.")
			return
		}

		context.Status(http.StatusOK)

	}
}

func (ServiceRequesterManagementService) GetStatistics() gin.HandlerFunc {
	return func(context *gin.Context) {
		startingDate := context.Query("startingDate")
		endingDate := context.Query("endingDate")
		requesterId, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The id you provided has a non-valid format.")
			return
		}

		_, parseError := time.Parse("2006-01-02", startingDate)
		if parseError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The starting date you provided has a non-valid format.")
			return
		}

		_, parseError = time.Parse("2006-01-02", endingDate)
		if parseError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ending date you provided has a non-valid format.")
			return
		}

		serviceRequester := businessEntities.ServiceRequester{
			ID: requesterId,
		}

		statisticsReport := struct {
			RequestedServicesPerWeekdayqueryResult []struct {
				RequestedServices int `json:"requestedServices"`
				Weekday           int `json:"weekday"`
			} `json:"requestedServicesPerWeekday"`

			KindOfServicesQueryResult []struct {
				RequestedServices int `json:"requestedServices"`
				KindOfService     int `json:"kindOfService"`
			} `json:"requestedServicesPerKindOfService"`
		}{}

		databaseError := serviceRequester.GetStatisticsReport(&statisticsReport.RequestedServicesPerWeekdayqueryResult, &statisticsReport.KindOfServicesQueryResult, startingDate, endingDate)
		if databaseError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to get the statistics")
			return
		}

		if statisticsReport.KindOfServicesQueryResult == nil || statisticsReport.RequestedServicesPerWeekdayqueryResult == nil {
			context.AbortWithStatusJSON(http.StatusNotFound, "No service requests were found for the dates or the user you provided.")
			return
		}
		context.JSON(http.StatusOK, statisticsReport)
	}
}
