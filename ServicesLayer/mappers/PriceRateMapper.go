package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects")


func CreatePriceRateDTOSliceAsResponse(priceRates []businessEntities.PriceRate) []dataTransferObjects.ResponsePriceRateDTO {

	var response []dataTransferObjects.ResponsePriceRateDTO
	for _, priceRateElement := range priceRates {
		priceRate := dataTransferObjects.ResponsePriceRateDTO{
			ID: priceRateElement.ID,
			StartingHour: priceRateElement.StartingHour,
			EndingHour: priceRateElement.EndingHour,
			Price: priceRateElement.Price,
			KindOfService: priceRateElement.KindOfService,
			City: CreateCityDTOAsResponse(priceRateElement.City),
			WorkingDays: CreateWorkingDayDTOSliceAsResponse(priceRateElement.WorkingDays),
		}
		response = append(response, priceRate)
	}


	return response
}