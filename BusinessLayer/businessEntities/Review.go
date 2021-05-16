package businessEntities

import (
	"time"

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
}