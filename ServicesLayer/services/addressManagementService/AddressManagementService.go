package addressManagementService

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	uuid "github.com/satori/go.uuid"
)


type AddressManagementService struct{}

func (AddressManagementService) Register(address businessEntities.Address) error {
	databaseError := address.Register()
	return databaseError
}

func (AddressManagementService) FindAll(serviceRequesterID uuid.UUID) ([]businessEntities.Address, error) {
	address := businessEntities.Address{}
	addresses, databaseError := address.FindAllAddressesOfServiceRequester(serviceRequesterID)
	return addresses, databaseError
	
}
