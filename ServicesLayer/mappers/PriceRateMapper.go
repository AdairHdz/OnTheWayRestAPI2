package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	uuid "github.com/satori/go.uuid"
)


func CreatePriceRateDTOSliceAsResponse(priceRates []businessEntities.PriceRate) []dataTransferObjects.ResponsePriceRateDTOWithCity {

	var response []dataTransferObjects.ResponsePriceRateDTOWithCity
	for _, priceRateElement := range priceRates {
		priceRate := dataTransferObjects.ResponsePriceRateDTOWithCity{
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

func CreatePriceRateDTOAsResponse(priceRate businessEntities.PriceRate) dataTransferObjects.ResponsePriceRateDTO {

	var response dataTransferObjects.ResponsePriceRateDTO

	response = dataTransferObjects.ResponsePriceRateDTO {
		ID: priceRate.ID,
		StartingHour: priceRate.StartingHour,
		EndingHour: priceRate.EndingHour,
		Price: priceRate.Price,
		KindOfService: priceRate.KindOfService,
		CityID: priceRate.CityID,		
		WorkingDays: CreateWorkingDayDTOSliceAsResponse(priceRate.WorkingDays),
	}
	
	
	return response
}

func CreatePriceRateEntity(priceRateDTO dataTransferObjects.ReceivedPriceRateDTO, serviceProviderID uuid.UUID) businessEntities.PriceRate {
	
	var workingDayEntities []businessEntities.WorkingDay

	for _, workingDayElement := range priceRateDTO.WorkingDays {
		workingDay := businessEntities.WorkingDay {
			ID: workingDayElement,
		}

		workingDayEntities = append(workingDayEntities, workingDay)
	}

	priceRate := businessEntities.PriceRate{
		ID: uuid.NewV4(),
		StartingHour: priceRateDTO.StartingHour,
		EndingHour: priceRateDTO.EndingHour,
		Price: priceRateDTO.Price,
		WorkingDays: workingDayEntities,
		ServiceProviderID: serviceProviderID,
		CityID: priceRateDTO.CityID,
		KindOfService: priceRateDTO.KindOfService,
	}

	return priceRate
}