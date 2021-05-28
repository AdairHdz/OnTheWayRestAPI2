package reviewManagementService

import (
	"fmt"
	"net/http"
	"path/filepath"

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
	return func(context *gin.Context){

		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}		

		receivedData := dataTransferObjects.ReceivedReviewDTO{}
		context.BindJSON(&receivedData)

		validate := validators.GetValidator()
		validationErrors := validate.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		review := mappers.CreateReviewEntity(receivedData, serviceProviderID)
		databaseError := review.Register()

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateResponseReviewDTO(review)

		context.JSON(http.StatusCreated, response)
	}
}

func (ReviewManagementService) Find() gin.HandlerFunc {
	return func(context *gin.Context){				
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		review := businessEntities.Review{}
		reviews, databaseError := review.Find(serviceProviderID)
		
		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := mappers.CreateSliceOfResponseReviewDTO(reviews)

		if len(response) == 0 {
			context.Status(http.StatusNotFound)
			return
		}
		context.JSON(http.StatusOK, response)		
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
			fileSavingError := context.SaveUploadedFile(file, path + "/" + file.Filename)
			if fileSavingError != nil {
				context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to save the evidence")
			}
		}

		context.Status(http.StatusCreated)
	}
}
