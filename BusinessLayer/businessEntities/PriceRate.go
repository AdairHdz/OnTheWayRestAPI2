package businessEntities

import (
	"time"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories/priceRateRepository"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type PriceRate struct {
	gorm.Model
	ID uuid.UUID
	StartingHour time.Time `gorm:"type:time not null;"`
	EndingHour time.Time `gorm:"type:time not null;"`
	Price float32
	WorkingDays []WorkingDay `gorm:"many2many:pricerate_workingday;constraint:OnDelete:CASCADE;"`
	ServiceProviderID uuid.UUID `gorm:"size:191"`
	CityID uuid.UUID `gorm:"size:191"`
	City City
	KindOfService uint8
}

func (priceRate PriceRate) Register() error {
	repository := repositories.Repository{}
	databaseError := repository.Create(&priceRate)
	return databaseError
}

func (priceRate PriceRate) Find(serviceProviderID uuid.UUID) ([]PriceRate, error) {
	repository := priceRateRepository.PriceRateRepository{}
	var priceRates []PriceRate
	databaseError :=  repository.FindMatches(&priceRates, "service_provider_id = ?", serviceProviderID)
	return priceRates, databaseError	
}

func (priceRate *PriceRate) Delete(serviceProviderID uuid.UUID) error {
	repository := repositories.Repository{}
	databaseError := repository.Delete(&priceRate, "service_provider_id = ?", serviceProviderID)
	return databaseError
}