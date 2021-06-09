package reviewManagementService

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/directoryManager"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/fileAnalyzer"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ReviewManagementService struct{}

func (ReviewManagementService) Register() gin.HandlerFunc {
	return func(context *gin.Context) {

		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format.")
			return
		}

		receivedData := dataTransferObjects.ReceivedReviewDTO{}
		context.BindJSON(&receivedData)

		validate := validators.GetValidator()
		validationErrors := validate.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
			return
		}

		review := mappers.CreateReviewEntity(receivedData, serviceProviderID)
		databaseError := review.Register()

		if databaseError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to register your review.")
			return
		}

		response := mappers.CreateResponseReviewDTO(review)

		context.JSON(http.StatusCreated, response)
	}
}

func (ReviewManagementService) Find() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format.")
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

		var rowCount int64
		review := businessEntities.Review{}
		reviews, databaseError := review.Find(page, pagesize, &rowCount, serviceProviderID)

		if databaseError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to retrieve the data you requested.")
			return
		}

		if len(reviews) == 0 {
			context.AbortWithStatusJSON(http.StatusNotFound, "There are no matches for the parameters you provided.")
			return
		}

		response := mappers.CreateSliceOfResponseReviewDTO(reviews)
		lastPage := int(rowCount / int64(pagesize))
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
				First string
				Last  string
				Prev  string
				Next  string
			}
			Page    int
			Pages   int
			PerPage int
			Total   int64
			Data    []dataTransferObjects.ResponseReviewDTOWithServiceRequesterData
		}{
			Links: struct {
				First string
				Last  string
				Prev  string
				Next  string
			}{
				First: fmt.Sprintf("providers/%s/reviews?page=%d&pagesize=%d", serviceProviderID, 1, pagesize),
				Last:  fmt.Sprintf("providers/%s/reviews?page=%d&pagesize=%d", serviceProviderID, lastPage, pagesize),
				Prev:  fmt.Sprintf("providers/%s/reviews?page=%d&pagesize=%d", serviceProviderID, previousPage, pagesize),
				Next:  fmt.Sprintf("providers/%s/reviews?page=%d&pagesize=%d", serviceProviderID, nextPage, pagesize),
			},
			Page:    page,
			Pages:   lastPage,
			PerPage: pagesize,
			Total:   rowCount,
			Data:    response,
		}
		context.JSON(http.StatusOK, dataResponse)
	}
}

func (ReviewManagementService) UploadEvidence() gin.HandlerFunc {
	return func(context *gin.Context) {
		var reviewId string = context.Param("reviewId")
		const maxFileSize = 10855731
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]
		path := fmt.Sprintf("./public/reviews/%s", reviewId)
		directoryCreationError := directoryManager.CreateDirectory(path)

		if directoryCreationError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save the evidence")
			return
		}

		dirIsEmpty, directoryEmptinessVerificationError := fileAnalyzer.DirIsEmpty(path)

		if directoryEmptinessVerificationError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save the evidence")
			return
		}

		if !dirIsEmpty {
			context.AbortWithStatusJSON(http.StatusConflict, "Attempted to add files to a review that already has files registered")
			return
		}

		if len(files) == 0 {
			context.AbortWithStatusJSON(http.StatusBadRequest, "Request should contain at least one file")
			return
		} else if len(files) > 3 {
			context.AbortWithStatusJSON(http.StatusBadRequest, "Can't upload more than 3 files per request")
			return
		}

		for _, file := range files {
			var fileSizeTotal int64 = file.Size
			if fileSizeTotal > maxFileSize {
				context.AbortWithStatusJSON(http.StatusConflict, "One or more files have a size greater than 10 MB")
				return
			}
			fileExtension := filepath.Ext(file.Filename)
			if !fileAnalyzer.EvidenceHasValidFormat(fileExtension) {
				context.AbortWithStatusJSON(http.StatusBadRequest, "One or more files have invalid format")
				return
			}
		}

		for _, file := range files {
			fileSavingError := context.SaveUploadedFile(file, path+"/"+file.Filename)
			if fileSavingError != nil {
				context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save the evidence")
			}
		}

		context.Status(http.StatusCreated)
	}
}
