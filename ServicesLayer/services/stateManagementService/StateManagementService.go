package stateManagementService

import (
	"net/http"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)


type StateManagementService struct{}

func (StateManagementService) FindAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		state := businessEntities.State{}
		states, databaseError := state.FindAll()

		if databaseError != nil {
			context.Status(http.StatusConflict)
			return
		}
		
		response := mappers.CreateSliceOfStateDTOAsResponse(states)

		context.JSON(http.StatusOK, response)
	}
}

func (StateManagementService) FindAllCitiesOfState() gin.HandlerFunc {
	return func(context *gin.Context) {
		stateID, parsingError := uuid.FromString(context.Param("stateId"))

		if parsingError != nil {
			context.Status(http.StatusConflict)
			return
		}

		city := businessEntities.City{}
		cities, databaseError := city.FindAll(stateID)

		if databaseError != nil {
			context.Status(http.StatusConflict)
			return
		}

		response := mappers.CreateSliceOfCityDTOAsResponse(cities)

		context.JSON(http.StatusOK, response)
	}
}