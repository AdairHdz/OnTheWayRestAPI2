package businessEntities

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories/serviceProviderRepository"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	gorm.Model
	ID uuid.UUID
	User User
	UserID uuid.UUID `gorm:"size:191"`
	AverageScore float32
	Reviews []Review
	PriceRates []PriceRate
}

func (serviceProvider ServiceProvider) Register() error {
	repository := repositories.Repository{}
	databaseError := repository.Create(&serviceProvider)
	return databaseError
}

func (serviceProvider *ServiceProvider) Find(serviceProviderID uuid.UUID) (error) {	
	repository := serviceProviderRepository.ServiceProviderRepository{}	
	databaseError := repository.FindByID(&serviceProvider, serviceProviderID)	
	return databaseError
	
}