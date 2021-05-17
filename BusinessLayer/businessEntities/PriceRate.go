package businessEntities

import (
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type PriceRate struct {
	gorm.Model
	ID uuid.UUID
	StartingHour time.Time
	EndingHour time.Time
	Price float32
	WorkingDays []WorkingDay `gorm:"many2many:pricerate_workingday;"`
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