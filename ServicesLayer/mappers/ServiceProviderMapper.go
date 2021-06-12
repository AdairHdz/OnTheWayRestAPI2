package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
)

func CreateServiceProviderDTOAsResponse(serviceProvider businessEntities.ServiceProvider) dataTransferObjects.ResponseServiceProviderDTO {
	response := dataTransferObjects.ResponseServiceProviderDTO{
		ID:           serviceProvider.ID,
		Names:        serviceProvider.User.Names,
		LastName:     serviceProvider.User.LastName,
		EmailAddress: serviceProvider.User.EmailAddress,
		AverageScore: uint8(serviceProvider.AverageScore),
		PriceRates:   CreatePriceRateDTOSliceAsResponse(serviceProvider.PriceRates),
	}
	return response
}

func CreateServiceProviderOverviewDTOAsResponse(serviceProviders []businessEntities.ServiceProvider, maxPriceRate float64, kindOfService uint8) []dataTransferObjects.ResponseServiceProviderOverviewDTO {
	var response []dataTransferObjects.ResponseServiceProviderOverviewDTO

	for _, serviceProviderElement := range serviceProviders {

		var foundPriceRate float64 = 0
		for _, priceRate := range serviceProviderElement.PriceRates {
			if float64(priceRate.Price) <= maxPriceRate && priceRate.KindOfService == kindOfService {
				foundPriceRate = float64(priceRate.Price)
				break
			}
		}

		if len(serviceProviderElement.PriceRates) != 0 {
			serviceProviderDTO := dataTransferObjects.ResponseServiceProviderOverviewDTO{
				ID:           serviceProviderElement.ID,
				Names:        serviceProviderElement.User.Names,
				LastName:     serviceProviderElement.User.LastName,
				AverageScore: uint8(serviceProviderElement.AverageScore),
				PriceRate:    float32(foundPriceRate),
			}
			response = append(response, serviceProviderDTO)
		}

	}

	return response
}
