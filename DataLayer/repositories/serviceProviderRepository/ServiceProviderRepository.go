package serviceProviderRepository

import "github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"

type ServiceProviderRepository struct{}

func (ServiceProviderRepository) FindByID(serviceProvider interface{}, ID interface{}) error {
	DB := database.GetDatabase()
	result := DB.Preload("User").Preload("User.State").First(serviceProvider, ID)
	return result.Error
}
