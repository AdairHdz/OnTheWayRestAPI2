package businessEntities

import (
	"time"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories/reviewRepository"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type Review struct {
	gorm.Model
	ID uuid.UUID
	DateOfReview time.Time
	Title string
	Details string
	Score uint8
	ServiceProviderID uuid.UUID `gorm:"size:191"`
	ServiceRequesterID uuid.UUID `gorm:"size:191"`
	ServiceRequester ServiceRequester
	Evidence []ReviewEvidence
}

func (review *Review) Register() error {
	repository := repositories.Repository{}
	databaseError := repository.Create(&review)
	return databaseError
}

func (review *Review) Find(serviceProviderID uuid.UUID) ([]Review, error) {	
	var reviews []Review
	repository := reviewRepository.ReviewRepository{}
	databaseError := repository.FindMatches(&reviews, "service_provider_id = ?", serviceProviderID)	
	return reviews, databaseError
	
}