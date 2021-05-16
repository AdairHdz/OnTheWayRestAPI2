package businessEntities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type ServiceRequest struct {
	gorm.Model
	ID uuid.UUID
	Cost float32
	Date time.Time
	AddressID uuid.UUID `gorm:"size:191"`
	DeliveryAddress Address `gorm:"foreignKey:AddressID"`
	Description string
	KindOfService uint8
	ServiceStatus uint8
	ServiceRequesterID uuid.UUID `gorm:"size:191"`
	ServiceProviderID uuid.UUID `gorm:"size:191"`
	ServiceRequester ServiceRequester
	ServiceProvider ServiceProvider
}