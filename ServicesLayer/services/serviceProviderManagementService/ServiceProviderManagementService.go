package serviceProviderManagementService

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/customErrors"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/directoryManager"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/fileAnalyzer"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ServiceProviderManagementService struct{}

func (ServiceProviderManagementService) Find() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		serviceProvider := businessEntities.ServiceProvider{}
		searchError := serviceProvider.Find(serviceProviderID)

		if searchError != nil {
			_, errorIsOfTypeRecordNotFoundError := searchError.(customErrors.RecordNotFoundError)
			if errorIsOfTypeRecordNotFoundError {
				context.AbortWithStatus(http.StatusNotFound)
				return
			}
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateServiceProviderDTOAsResponse(serviceProvider)

		context.JSON(http.StatusOK, response)
	}
}

func (ServiceProviderManagementService) FindMatches() gin.HandlerFunc {
	return func(context *gin.Context) {
		maxPriceRate, parseError := strconv.ParseFloat(context.Query("maxPriceRate"), 32)

		if parseError != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		city := context.Query("city")
		kindOfService, parseError := strconv.ParseInt(context.Query("kindOfService"), 10, 8)
		serviceProvider := businessEntities.ServiceProvider{}
		serviceProviders, err := serviceProvider.FindMatches(maxPriceRate, city, kindOfService)

		if err != nil {
			context.Status(http.StatusConflict)
			return
		}

		if len(serviceProviders) == 0 {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}
		response := mappers.CreateServiceProviderOverviewDTOAsResponse(serviceProviders)
		context.JSON(http.StatusOK, response)
	}
}

func (ServiceProviderManagementService) Update() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
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

		serviceProvider := businessEntities.ServiceProvider{}
		serviceProvider.Find(serviceProviderID)

		if receivedData.Names != "" {
			serviceProvider.User.Names = receivedData.Names
		}

		if receivedData.LastName != "" {
			serviceProvider.User.LastName = receivedData.LastName
		}

		if receivedData.Password != "" {
			hashedPassword, hashingError := hashing.GenerateHash(serviceProvider.User.Password)
			if hashingError != nil {
				context.AbortWithStatus(http.StatusConflict)
				return
			}

			serviceProvider.User.Password = hashedPassword
		}

		validator := validators.GetValidator()
		validationErrors := validator.Var(serviceProvider.User.Names, "required,min=1,max=50,lettersAndSpaces")

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		validationErrors = validator.Var(serviceProvider.User.LastName, "required,min=1,max=50,lettersAndSpaces")

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		validationErrors = validator.Var(serviceProvider.User.Password, "required,max=80")

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		updateError := serviceProvider.Update()

		if updateError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		context.Status(http.StatusNoContent)
	}
}

func (ServiceProviderManagementService) UpdateServiceProviderImage() gin.HandlerFunc {
	return func(context *gin.Context) {
		providerID := context.Param("providerId")
		path := "./images/" + providerID
		directoryCreationError := directoryManager.CreateDirectory(path)

		if directoryCreationError != nil {
			context.AbortWithStatus(http.StatusConflict)
		}

		serviceProvider := businessEntities.ServiceProvider{}

		serviceProvider.Find(uuid.FromStringOrNil(providerID))

		if serviceProvider.ID == uuid.Nil {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		dirIsEmpty, err := fileAnalyzer.DirIsEmpty(path)

		file, noFileSentError := context.FormFile("image")
		if noFileSentError != nil {
			context.Status(http.StatusBadRequest)
			return
		}

		fileExtension := filepath.Ext(file.Filename)

		if !fileAnalyzer.ImageHasValidFormat(fileExtension) {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		if !dirIsEmpty {
			pathOfImageToBeDeleted := path + "/" + serviceProvider.BusinessPicture
			os.Remove(pathOfImageToBeDeleted)
		}

		err = context.SaveUploadedFile(file, path+"/"+file.Filename)

		if err != nil {
			context.JSON(http.StatusConflict, err.Error())
			return
		}

		serviceProvider.BusinessPicture = file.Filename

		databaseError := serviceProvider.Update()

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		context.Status(http.StatusOK)
	}
}

func (ServiceProviderManagementService) GetStatistics() gin.HandlerFunc {
	return func(context *gin.Context) {
		startingDate := context.Query("startingDate")
		endingDate := context.Query("endingDate")
		providerId, parsingError := uuid.FromString(context.Param("providerId"))

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

		serviceProvider := businessEntities.ServiceProvider{
			ID: providerId,
		}

		statisticsReport := struct {
			RequestedServicesPerWeekdayqueryResult []struct {
				RequestedServices int `json:"requestedServices"`
				Weekday           int `json:"weekday"`
			}

			KindOfServicesQueryResult []struct {
				RequestedServices int `json:"requestedServices"`
				KindOfService     int `json:"kindOfService"`
			}
		}{}

		databaseError := serviceProvider.GetStatisticsReport(&statisticsReport.RequestedServicesPerWeekdayqueryResult, &statisticsReport.KindOfServicesQueryResult, startingDate, endingDate)
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
