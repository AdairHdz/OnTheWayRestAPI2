package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"	
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
)


func CreateServiceProviderDTOAsResponse(serviceProvider businessEntities.ServiceProvider) dataTransferObjects.ResponseServiceProviderDTO {
	response := dataTransferObjects.ResponseServiceProviderDTO{
		ID: serviceProvider.ID,
		Names: serviceProvider.User.Names,
		LastName: serviceProvider.User.LastName,
		EmailAddress: serviceProvider.User.EmailAddress,
		AverageScore: uint8(serviceProvider.AverageScore),
		PriceRates: CreatePriceRateDTOSliceAsResponse(serviceProvider.PriceRates),
	}
	return response
}

func CreateServiceProviderOverviewDTOAsResponse(serviceProviders []businessEntities.ServiceProvider) []dataTransferObjects.ResponseServiceProviderOverviewDTO {
	var response []dataTransferObjects.ResponseServiceProviderOverviewDTO	
	
	for _, serviceProviderElement := range serviceProviders {

		if len(serviceProviderElement.PriceRates) != 0 {			
			serviceProviderDTO := dataTransferObjects.ResponseServiceProviderOverviewDTO {
				ID: serviceProviderElement.ID,
				Names: serviceProviderElement.User.Names,
				LastName: serviceProviderElement.User.LastName,
				AverageScore: uint8(serviceProviderElement.AverageScore),
				PriceRate: serviceProviderElement.PriceRates[0].Price,
			}
			response = append(response, serviceProviderDTO)
		}


	}

	
	return response
}