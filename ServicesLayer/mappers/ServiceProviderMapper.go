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