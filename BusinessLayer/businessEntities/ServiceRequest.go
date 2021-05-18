package businessEntities

import (
	"time"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories/serviceRequestRepository"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const (
	PendingOfAcceptance = iota
	Active = iota
	Concluded = iota
	Canceled = iota
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

func (serviceRequest *ServiceRequest) Register() error {
	repository := repositories.Repository{}
	result := repository.Create(&serviceRequest)
	return result
}

func (serviceRequest *ServiceRequest) Find(serviceRequestId uuid.UUID) error {	
	repository := serviceRequestRepository.ServiceRequestRepository{}
	databaseError := repository.FindByID(&serviceRequest, serviceRequestId)
	return databaseError
}


func (serviceRequest *ServiceRequest) Update() error{
	repository := repositories.Repository{}	
	databaseError := repository.Update(&serviceRequest)
	return databaseError
}