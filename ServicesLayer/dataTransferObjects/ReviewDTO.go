package dataTransferObjects

import (
	"time"
	uuid "github.com/satori/go.uuid")


type ReceivedReviewDTO struct {
	Title              string `json:"title"`
	Details            string `json:"details"`
	Score              uint8  `json:"score"`
	Evidence []ReviewEvidenceDTO `json:"evidence"`
	ServiceRequesterID uuid.UUID `json:"serviceRequesterId"`
}

type ResponseReviewDTO struct {
	ID uuid.UUID `json:"id"`
	DateOfReview time.Time `json:"dateOfReview"`
	Title              string `json:"title"`
	Details            string `json:"details"`
	Score              uint8  `json:"score"`
	Evidence []ReviewEvidenceDTO `json:"evidence"`
	ServiceRequesterID uuid.UUID `json:"serviceRequesterId"`
}


type ResponseReviewDTOWithServiceRequesterData struct {
	ID uuid.UUID `json:"id"`
	DateOfReview time.Time `json:"dateOfReview"`
	Title              string `json:"title"`
	Details            string `json:"details"`
	Score              uint8  `json:"score"`
	Evidence []ReviewEvidenceDTO `json:"evidence"`
	ServiceRequester struct {
		ID uuid.UUID `json:"id"`
		Name string `json:"name"`
		LastName string `json:"lastName"`
	} `json:"serviceRequester"`
}