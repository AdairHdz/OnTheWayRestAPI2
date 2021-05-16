package businessEntities

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories/addressRepository"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID uuid.UUID
	IndoorNumber  string
	OutdoorNumber string
	Street        string
	Suburb        string
	ServiceRequesterID uuid.UUID `gorm:"size:191"`
	CityID uuid.UUID `gorm:"size:191"`
	City City
}

func (address Address) Register() error {
	repository := repositories.Repository{}
	databaseError := repository.Create(&address)
	return databaseError
}

func (Address) FindAllAddressesOfServiceRequester(serviceRequesterID uuid.UUID) ([]Address, error) {
	repository := addressRepository.AddressRepository{}
	var addresses []Address
	databaseError :=  repository.FindMatches(&addresses, "service_requester_id = ?", serviceRequesterID)
	return addresses, databaseError	
}