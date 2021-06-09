package stateManagementService

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/cacheManager"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
)

type StateManagementService struct{}

func (StateManagementService) FindAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		cachedStatesBytes, cacheError := cacheManager.Get("states")
		if cacheError == nil && cachedStatesBytes != nil {
			var cachedStates []dataTransferObjects.StateDTO
			unmarshallingError := json.Unmarshal(cachedStatesBytes, &cachedStates)
			if unmarshallingError == nil {
				context.JSON(http.StatusOK, cachedStates)
				return
			}
		}

		state := businessEntities.State{}
		states, databaseError := state.FindAll()

		if databaseError != nil {
			context.Status(http.StatusConflict)
			return
		}

		response := mappers.CreateSliceOfStateDTOAsResponse(states)

		if cacheError == redis.Nil {
			jsonStates, jsonError := json.Marshal(response)
			if jsonError == nil {
				cacheManager.Save("states", jsonStates, time.Hour*2)
			}
		}
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

		cachedCitiesBytes, cacheError := cacheManager.Get(stateID.String())
		if cacheError == nil && cachedCitiesBytes != nil {
			var cachedCities []dataTransferObjects.CityDTO
			unmarshallingError := json.Unmarshal(cachedCitiesBytes, &cachedCities)
			if unmarshallingError == nil {
				context.JSON(http.StatusOK, cachedCities)
				return
			}
		}

		city := businessEntities.City{}
		cities, databaseError := city.FindAll(stateID)

		if databaseError != nil {
			context.Status(http.StatusConflict)
			return
		}

		response := mappers.CreateSliceOfCityDTOAsResponse(cities)

		if cacheError == redis.Nil {
			jsonStates, jsonError := json.Marshal(response)
			if jsonError == nil {
				cacheManager.Save(stateID.String(), jsonStates, time.Hour*2)
			}
		}

		context.JSON(http.StatusOK, response)
	}
}
