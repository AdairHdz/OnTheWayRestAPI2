package businessEntities

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories/serviceProviderRepository"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	gorm.Model
	ID               uuid.UUID
	User             User
	UserID           uuid.UUID `gorm:"size:191"`
	AverageScore     float32
	TotalPoints      int
	MaxTotalPossible int
	Reviews          []Review
	PriceRates       []PriceRate
	BusinessPicture  string
}

func (serviceProvider ServiceProvider) Register() error {
	repository := repositories.Repository{}
	databaseError := repository.Create(&serviceProvider)
	return databaseError
}

func (serviceProvider *ServiceProvider) Find(serviceProviderID uuid.UUID) error {
	repository := serviceProviderRepository.ServiceProviderRepository{}
	databaseError := repository.FindByID(&serviceProvider, serviceProviderID)
	return databaseError

}

func (serviceProvider *ServiceProvider) Update() error {
	repository := serviceProviderRepository.ServiceProviderRepository{}
	databaseError := repository.Update(&serviceProvider.User)
	if databaseError != nil {
		return databaseError
	}

	databaseError = repository.Update(&serviceProvider)
	return databaseError
}

func (ServiceProvider) FindMatches(page, pagesize int, numberOfRows *int64, maxPriceRate float64, cityName string, kindOfService int64) ([]ServiceProvider, error) {
	var serviceProviders []ServiceProvider
	repository := serviceProviderRepository.ServiceProviderRepository{}
	err := repository.FindMatches(&serviceProviders, page, pagesize, numberOfRows, maxPriceRate, cityName, kindOfService)
	return serviceProviders, err
}

func (ServiceProvider *ServiceProvider) GetStatisticsReport(requestedServicesPerWeekdayqueryResult, kindOfServicesQueryResult interface{}, startingDate, endingDate string) error {
	repository := serviceProviderRepository.ServiceProviderRepository{}
	databaseError := repository.GetStatisticsReport(requestedServicesPerWeekdayqueryResult, kindOfServicesQueryResult, ServiceProvider.ID.String(), startingDate, endingDate)
	return databaseError
}
