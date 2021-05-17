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