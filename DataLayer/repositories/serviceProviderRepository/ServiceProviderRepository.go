package serviceProviderRepository

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database")


type ServiceProviderRepository struct{}

func (ServiceProviderRepository) FindByID(serviceProvider interface{}, ID interface{}) error {
	DB := database.GetDatabase()
	result := DB.Preload("User").Preload("Reviews").Preload("PriceRates.WorkingDays").Preload("PriceRates.City").First(serviceProvider, ID)
	return result.Error
}

func (ServiceProviderRepository) Update(serviceProvider interface{}) (error) {
	DB := database.GetDatabase()
	result := DB.Updates(serviceProvider)
	return result.Error
}

func (ServiceProviderRepository) FindMatches(target interface{}, maxPriceRate float64, cityName string, kindOfService int64) (error) {	
	DB := database.GetDatabase()
	result := DB.Preload("User").Preload("PriceRates").
	Joins("INNER JOIN users ON users.id = service_providers.user_id").
	Joins("INNER JOIN price_rates ON price_rates.service_provider_id = service_providers.id").
	Joins("INNER JOIN cities ON price_rates.city_id = cities.id").
	Where("price_rates.price <= ? AND cities.name = ? AND price_rates.kind_of_service = ?", maxPriceRate, cityName, kindOfService).
	Find(target)	
	return result.Error
}