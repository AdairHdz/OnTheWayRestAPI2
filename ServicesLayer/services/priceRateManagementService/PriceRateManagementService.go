package priceRateManagementService

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/mappers"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/customErrors"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type PriceRateManagementService struct{}

func (PriceRateManagementService) Register() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format.")
			return
		}

		receivedData := dataTransferObjects.ReceivedPriceRateDTO{}

		context.BindJSON(&receivedData)

		validator := validators.GetValidator()
		validationErrors := validator.Struct(receivedData)

		if validationErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
			return
		}

		priceRateEntity, mappingError := mappers.CreatePriceRateEntity(receivedData, serviceProviderID)

		if mappingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
			return
		}

		databaseError := priceRateEntity.Register()

		if databaseError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to register the resource.")
			return
		}

		response := mappers.CreatePriceRateDTOAsResponse(priceRateEntity)
		context.JSON(http.StatusCreated, response)
	}
}

func (PriceRateManagementService) FindAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format.")
			return
		}

		kindOfService := context.Query("kindOfService")
		city := context.Query("city")
		priceRate := businessEntities.PriceRate{}
		priceRates, databaseError := priceRate.Find(serviceProviderID)

		if databaseError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to retrieve the data.")
			return
		}

		response := mappers.CreatePriceRateDTOSliceAsResponse(priceRates)
		if len(response) == 0 {
			context.AbortWithStatusJSON(http.StatusNotFound, "There are no price rates registered for this service provider.")
			return
		}

		if kindOfService != "" || city != "" {
			validator := validators.GetValidator()
			validationErrors := validator.Var(kindOfService, "min=0,max=4")
			if validationErrors != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
				return
			}

			validationErrors = validator.Var(city, "min=1,max=50,lettersAndSpaces")
			if validationErrors != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
				return
			}

			location, locationLoadingError := time.LoadLocation("America/Mexico_City")
			if locationLoadingError != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, "There was an eror while trying to retrieve the active price rate.")
				return
			}
			parsedCurrentTime, parsingError := time.Parse(time.Kitchen, time.Now().In(location).Format(time.Kitchen))
			if parsingError != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, "There was an eror while trying to retrieve the price rates.")
				return
			}

			activePriceRate := dataTransferObjects.ResponsePriceRateDTOWithCity{}

			for _, priceRateElement := range response {
				parsedStartingTime, parsingError := time.Parse(time.Kitchen, priceRateElement.StartingHour)
				if parsingError != nil {
					context.AbortWithStatusJSON(http.StatusBadRequest, "There was an eror while trying to retrieve the active price rate.")
					return
				}

				parsedEndingTime, parsingError := time.Parse(time.Kitchen, priceRateElement.EndingHour)
				if parsingError != nil {
					context.AbortWithStatusJSON(http.StatusBadRequest, "There was an eror while trying to retrieve the active price rate.")
					return
				}

				parsedkindOfService, parsingError := strconv.Atoi(kindOfService)

				if parsingError != nil {
					context.AbortWithStatusJSON(http.StatusBadRequest, "There was an eror while trying to retrieve the active price rate.")
					return
				}

				priceRateAppliesToCurrentDay := false

				for _, workingDayElement := range priceRateElement.WorkingDays {
					if workingDayElement == parsedCurrentTime.Day() {
						priceRateAppliesToCurrentDay = true
						break
					}
				}

				if parsedEndingTime.Sub(parsedStartingTime) < 0 {
					if parsedCurrentTime.Before(parsedEndingTime) &&
						priceRateElement.KindOfService == uint8(parsedkindOfService) &&
						priceRateElement.City.Name == city &&
						priceRateAppliesToCurrentDay {
						context.JSON(http.StatusOK, priceRateElement)
						return
					}
				} else {
					if parsedCurrentTime.After(parsedStartingTime) && parsedCurrentTime.Before(parsedEndingTime) &&
						priceRateElement.KindOfService == uint8(parsedkindOfService) &&
						priceRateElement.City.Name == city &&
						priceRateAppliesToCurrentDay {
						context.JSON(http.StatusOK, priceRateElement)
						return
					}
				}

			}

			if activePriceRate.ID == uuid.Nil {
				context.AbortWithStatusJSON(http.StatusNotFound, "There is not an active price rate for the current time.")
				return
			}
		} else {
			context.JSON(http.StatusOK, response)
		}

	}
}

func (PriceRateManagementService) Delete() gin.HandlerFunc {
	return func(context *gin.Context) {
		serviceProviderID, parsingError := uuid.FromString(context.Param("providerId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format.")
			return
		}

		priceRateID, parsingError := uuid.FromString(context.Param("priceRateId"))

		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The ID you provided has a non-valid format.")
			return
		}

		priceRate := businessEntities.PriceRate{
			ID: priceRateID,
		}

		databaseError := priceRate.Delete(serviceProviderID)

		if databaseError != nil {
			_, errorIsOfTypeRecordNotFound := databaseError.(customErrors.RecordNotFoundError)
			if errorIsOfTypeRecordNotFound {
				context.AbortWithStatusJSON(http.StatusNotFound, "There are no matches for the parameters of the resource you tried to delete.")
				return
			}
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to delete the resource.")
			return
		}

		context.Status(http.StatusNoContent)

	}
}
