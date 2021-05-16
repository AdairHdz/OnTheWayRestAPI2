package serviceRequesterManagementService

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	uuid "github.com/satori/go.uuid"
)


type ServiceRequesterManagementService struct{}

func (ServiceRequesterManagementService) Register(serviceRequester businessEntities.ServiceRequester) error {
	registryError := serviceRequester.Register()
	return registryError
}

func (ServiceRequesterManagementService) Find(serviceRequesterID uuid.UUID) (businessEntities.ServiceRequester,error) {
	var serviceRequester businessEntities.ServiceRequester
	searchError := serviceRequester.Find(serviceRequesterID)
	return serviceRequester, searchError
}

func (ServiceRequesterManagementService) Update(serviceRequester businessEntities.ServiceRequester) error {	
	updateError := serviceRequester.Update()
	return updateError
}
