package serviceProviderRepository

import "github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"

type ServiceProviderRepository struct{}

func (ServiceProviderRepository) FindByID(serviceProvider interface{}, ID interface{}) error {
	DB := database.GetDatabase()
	result := DB.Preload("User").Preload("Reviews").Preload("PriceRates").First(serviceProvider, ID)
	return result.Error
}

func (ServiceProviderRepository) Update(serviceProvider interface{}) (error) {
	DB := database.GetDatabase()
	result := DB.Updates(serviceProvider)
	return result.Error
}