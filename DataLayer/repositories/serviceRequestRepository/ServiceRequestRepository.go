package serviceRequestRepository

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/customErrors"
	uuid "github.com/satori/go.uuid"
)


type ServiceRequestRepository struct{}

func (ServiceRequestRepository) FindByID(serviceRequester interface{}, ID interface{}) error {
	DB := database.GetDatabase()
	result := DB.Preload("DeliveryAddress.City").Preload("ServiceProvider.User").Preload("ServiceRequester.User").First(serviceRequester, ID)
	if result.RowsAffected == 0 {
		return customErrors.RecordNotFoundError{}
	}
	return result.Error
}

func (ServiceRequestRepository) FindByDateAndServiceProviderID(target interface{}, date string, id uuid.UUID) (error) {
	DB := database.GetDatabase()
	result := DB.Preload("DeliveryAddress.City").Preload("ServiceProvider.User").Preload("ServiceRequester.User").Find(target, "date = ? AND service_provider_id = ?", date, id)
	return result.Error
}

func (ServiceRequestRepository) FindByDateAndServiceRequesterID(target interface{}, date string, id uuid.UUID) (error) {
	DB := database.GetDatabase()
	result := DB.Preload("DeliveryAddress.City").Preload("ServiceProvider.User").Preload("ServiceRequester.User").Find(target, "date = ? AND service_requester_id = ?", date, id)
	return result.Error
}