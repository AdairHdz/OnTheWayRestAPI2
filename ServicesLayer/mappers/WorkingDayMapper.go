package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"		
)


func CreateWorkingDayDTOSliceAsResponse(workingDays []businessEntities.WorkingDay) []int {
	
	var response []int

	for _, workingDayElement := range workingDays {		

		
		response = append(response, workingDayElement.ID)
	}

	return response
}