package mappers

import (
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	uuid "github.com/satori/go.uuid"
)


func CreateReviewEntity(reviewDTO dataTransferObjects.ReceivedReviewDTO, serviceProviderID uuid.UUID) businessEntities.Review {
	response := businessEntities.Review{
		ID: uuid.NewV4(),
		DateOfReview: time.Now(),
		Title: reviewDTO.Title,
		Details: reviewDTO.Details,
		Score: reviewDTO.Score,
		ServiceRequesterID: reviewDTO.ServiceRequesterID,
		ServiceProviderID: serviceProviderID,		
		Evidence: CreateSliceOfReviewEvidenceEntities(reviewDTO.Evidence),
	}
	return response
}

func CreateResponseReviewDTO(review businessEntities.Review) dataTransferObjects.ResponseReviewDTO {
	response := dataTransferObjects.ResponseReviewDTO {
		ID: review.ID,
		DateOfReview: review.DateOfReview,
		Title: review.Title,
		Details: review.Details,
		Score: review.Score,
		Evidence: CreateSliceOfReviewEvidenceDTOAsResponse(review.Evidence),
		ServiceRequesterID: review.ServiceRequesterID,
	}
	return response
}

func CreateSliceOfResponseReviewDTO(reviews []businessEntities.Review) []dataTransferObjects.ResponseReviewDTOWithServiceRequesterData {
	
	var response []dataTransferObjects.ResponseReviewDTOWithServiceRequesterData

	for _, reviewElement := range reviews {

		review := dataTransferObjects.ResponseReviewDTOWithServiceRequesterData {
			ID: reviewElement.ID,
			DateOfReview: reviewElement.DateOfReview,
			Title: reviewElement.Title,
			Details: reviewElement.Details,
			Score: reviewElement.Score,
			Evidence: CreateSliceOfReviewEvidenceDTOAsResponse(reviewElement.Evidence),
			ServiceRequester: struct{
				ID uuid.UUID "json:\"id\""; 
				Name string "json:\"name\""; 
				LastName string "json:\"lastName\""
			}{
				ID: reviewElement.ServiceRequester.ID,
				Name: reviewElement.ServiceRequester.User.Names,
				LastName: reviewElement.ServiceRequester.User.LastName,
			},
		}

		response = append(response, review)
	}

	return response
}