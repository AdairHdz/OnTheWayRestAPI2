package serviceRequestRepository

import "github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"

type ServiceRequestRepository struct{}

func (ServiceRequestRepository) FindByID(serviceRequester interface{}, ID interface{}) error {
	DB := database.GetDatabase()
	result := DB.Preload("DeliveryAddress.City").Preload("ServiceProvider.User").Preload("ServiceRequester.User").First(serviceRequester, ID)
	return result.Error
}