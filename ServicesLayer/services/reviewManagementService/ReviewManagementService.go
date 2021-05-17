package reviewManagementService

import (
	"net/http"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
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
		context.JSON(http.StatusOK, response)		
	}
}
