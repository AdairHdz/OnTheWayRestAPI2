package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"	
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
)


func CreateCityDTOAsResponse(city businessEntities.City) dataTransferObjects.CityDTO {
	response := dataTransferObjects.CityDTO {
		ID: city.ID,
		Name: city.Name,
	}
	return response
}

func CreateSliceOfCityDTOAsResponse(cities []businessEntities.City) []dataTransferObjects.CityDTO {
	
	var response []dataTransferObjects.CityDTO

	for _, cityElement := range cities {
		cityDTO := dataTransferObjects.CityDTO {
			ID: cityElement.ID,
			Name: cityElement.Name,
		}

		response = append(response, cityDTO)
		
	}
	
	return response
}