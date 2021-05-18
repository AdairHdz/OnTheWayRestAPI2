package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"	
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
)


func CreateSliceOfStateDTOAsResponse(states []businessEntities.State) []dataTransferObjects.StateDTO {
	var response []dataTransferObjects.StateDTO

	for _, stateElement := range states {
		stateDTO := dataTransferObjects.StateDTO {
			ID: stateElement.ID,
			Name: stateElement.Name,
		}

		response = append(response, stateDTO)
	}

	return response
}