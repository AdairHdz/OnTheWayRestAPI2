package dataTransferObjects

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ReceivedReviewDTO struct {
	Title              string              `json:"title" validate:"required,max=50"`
	Details            string              `json:"details" validate:"max=250"`
	Score              uint8               `json:"score" validate:"min=1,max=5"`
	Evidence           []ReviewEvidenceDTO `json:"evidence"`
	ServiceRequesterID uuid.UUID           `json:"serviceRequesterId"`
}

type ResponseReviewDTO struct {
	ID                 uuid.UUID                   `json:"id"`
	DateOfReview       time.Time                   `json:"dateOfReview"`
	Title              string                      `json:"title"`
	Details            string                      `json:"details"`
	Score              uint8                       `json:"score"`
	Evidence           []ReviewEvidenceRespondeDTO `json:"evidence"`
	ServiceRequesterID uuid.UUID                   `json:"serviceRequesterId"`
}

type ResponseReviewDTOWithServiceRequesterData struct {
	ID               uuid.UUID                    `json:"id"`
	DateOfReview     time.Time                    `json:"dateOfReview"`
	Title            string                       `json:"title"`
	Details          string                       `json:"details"`
	Score            uint8                        `json:"score"`
	Evidence         []ReviewEvidenceRespondeDTO  `json:"evidence"`
	ServiceRequester ResponseUserDTOWithNamesOnly `json:"serviceRequester"`
}
