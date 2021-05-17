package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"		
)


func CreateWorkingDayDTOSliceAsResponse(workingDays []businessEntities.WorkingDay) []uint8 {
	
	var response []uint8

	for _, workingDayElement := range workingDays {		

		response = append(response, workingDayElement.ID)
	}

	return response
}