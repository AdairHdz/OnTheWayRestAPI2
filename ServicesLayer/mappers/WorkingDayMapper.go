package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"	
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
)


func CreateWorkingDayDTOSliceAsResponse(workingDays []businessEntities.WorkingDay) []dataTransferObjects.ResponseWorkingDayDTO {
	
	var response []dataTransferObjects.ResponseWorkingDayDTO


	for _, workingDayElement := range workingDays {
		workingDay := dataTransferObjects.ResponseWorkingDayDTO {
			ID: workingDayElement.ID,
			Name: workingDayElement.Name,
		}

		response = append(response, workingDay)
	}

	return response
}