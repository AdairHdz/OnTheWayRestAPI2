package serviceProviderManagementService

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
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
			context.AbortWithStatusJSON(http.StatusBadRequest, "The Id you provided has a non-valid format.")
			return
		}

		serviceProvider := businessEntities.ServiceProvider{}
		searchError := serviceProvider.Find(serviceProviderID)

		if searchError != nil {
			_, errorIsOfTypeRecordNotFoundError := searchError.(customErrors.RecordNotFoundError)
			if errorIsOfTypeRecordNotFoundError {
				context.AbortWithStatusJSON(http.StatusNotFound, "There are no matches for the search parameters you provided.")
				return
			}
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to retrieve the requested data.")
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
			context.AbortWithStatusJSON(http.StatusBadRequest, "Invalid max price rate parameter")
			return
		}

		city := context.Query("city")
		kindOfService, parseError := strconv.ParseInt(context.Query("kindOfService"), 10, 8)
		if parseError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "Invalid kind of service parameter")
			return
		}

		page, conversionError := strconv.Atoi(context.Query("page"))
		if conversionError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "Invalid page parameter")
			return
		}

		pagesize, conversionError := strconv.Atoi(context.Query("pagesize"))
		if conversionError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "Invalid pagesize parameter")
			return
		}

		serviceProvider := businessEntities.ServiceProvider{}
		var count int64
		serviceProviders, err := serviceProvider.FindMatches(page, pagesize, &count, maxPriceRate, city, kindOfService)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to retrieve the data.")
			return
		}

		if len(serviceProviders) == 0 {
			context.AbortWithStatusJSON(http.StatusNotFound, "There are no matches for the data you requested.")
			return
		}

		response := mappers.CreateServiceProviderOverviewDTOAsResponse(serviceProviders, maxPriceRate, uint8(kindOfService))
		lastPage := (int(count) / pagesize)
		var previousPage int = 1
		var nextPage int

		if page-1 > 1 {
			previousPage = page - 1
		}

		if page+1 <= lastPage {
			nextPage = page + 1
		} else {
			nextPage = lastPage
		}

		dataResponse := struct {
			Links struct {
				First string `json:"first"`
				Last  string `json:"last"`
				Prev  string `json:"prev"`
				Next  string `json:"next"`
			} `json:"links"`
			Page    int                                                      `json:"page"`
			Pages   int                                                      `json:"pages"`
			PerPage int                                                      `json:"perPage"`
			Total   int64                                                    `json:"total"`
			Data    []dataTransferObjects.ResponseServiceProviderOverviewDTO `json:"data"`
		}{
			Links: struct {
				First string `json:"first"`
				Last  string `json:"last"`
				Prev  string `json:"prev"`
				Next  string `json:"next"`
			}{
				First: fmt.Sprintf("providers?maxPriceRate=%.2f&kindOfService=%d&city=%s&page=%d&pagesize=%d", maxPriceRate, kindOfService, city, 1, pagesize),
				Last:  fmt.Sprintf("providers?maxPriceRate=%.2f&kindOfService=%d&city=%s&page=%d&pagesize=%d", maxPriceRate, kindOfService, city, lastPage, pagesize),
				Prev:  fmt.Sprintf("providers?maxPriceRate=%.2f&kindOfService=%d&city=%s&page=%d&pagesize=%d", maxPriceRate, kindOfService, city, previousPage, pagesize),
				Next:  fmt.Sprintf("providers?maxPriceRate=%.2f&kindOfService=%d&city=%s&page=%d&pagesize=%d", maxPriceRate, kindOfService, city, nextPage, pagesize),
			},
			Page:    page,
			Pages:   lastPage,
			PerPage: pagesize,
			Total:   count,
			Data:    response,
		}
		context.JSON(http.StatusOK, dataResponse)
	}
}

func (ServiceProviderManagementService) Update() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format")
			return
		}

		receivedData := struct {
			Names    string `json:"names"`
			LastName string `json:"lastName"`
			Password string `json:"password"`
		}{}

		bindingError := context.BindJSON(&receivedData)

		if bindingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format")
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
				context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to update the resource")
				return
			}

			serviceProvider.User.Password = hashedPassword
		}

		validator := validators.GetValidator()
		validationErrors := validator.Var(serviceProvider.User.Names, "required,min=1,max=50,lettersAndSpaces")

		if validationErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format")
			return
		}

		validationErrors = validator.Var(serviceProvider.User.LastName, "required,min=1,max=50,lettersAndSpaces")

		if validationErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format")
			return
		}

		validationErrors = validator.Var(serviceProvider.User.Password, "required,max=80")

		if validationErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format")
			return
		}

		updateError := serviceProvider.Update()

		if updateError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to update the resource")
			return
		}

		context.Status(http.StatusOK)
	}
}

func (ServiceProviderManagementService) UpdateServiceProviderImage() gin.HandlerFunc {
	return func(context *gin.Context) {
		providerID := context.Param("providerId")
		path := "./images/" + providerID
		directoryCreationError := directoryManager.CreateDirectory(path)

		if directoryCreationError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save your image.")
			return
		}

		serviceProvider := businessEntities.ServiceProvider{}

		serviceProvider.Find(uuid.FromStringOrNil(providerID))

		if serviceProvider.ID == uuid.Nil {
			context.AbortWithStatusJSON(http.StatusNotFound, "There is not a service provider with the ID you provided.")
			return
		}

		dirIsEmpty, err := fileAnalyzer.DirIsEmpty(path)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save your image.")
			return
		}

		file, noFileSentError := context.FormFile("image")
		if noFileSentError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "You didn't provide any file.")
			return
		}

		fileExtension := filepath.Ext(file.Filename)

		if !fileAnalyzer.ImageHasValidFormat(fileExtension) {
			context.AbortWithStatusJSON(http.StatusConflict, "Invalid image format. Please make sure your file has jpg, jpeg, or png extension")
			return
		}

		if !dirIsEmpty {
			pathOfImageToBeDeleted := path + "/" + serviceProvider.BusinessPicture
			os.Remove(pathOfImageToBeDeleted)
		}

		err = context.SaveUploadedFile(file, path+"/"+file.Filename)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save your image.")
			return
		}

		serviceProvider.BusinessPicture = file.Filename

		databaseError := serviceProvider.Update()

		if databaseError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save your image.")
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
			} `json:"requestedServicesPerWeekday"`

			KindOfServicesQueryResult []struct {
				RequestedServices int `json:"requestedServices"`
				KindOfService     int `json:"kindOfService"`
			} `json:"requestedServicesPerKindOfService"`
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
