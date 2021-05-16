package addressController

import (
	"net/http"
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/addressManagementService"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
)


func RegisterAddress() gin.HandlerFunc{
	return func(context *gin.Context){

		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		receivedData := struct {
			IndoorNumber string `json:"indoorNumber" validate:"max=8"`
			OutdoorNumber string `json:"outdoorNumber" validate:"required,max=8"`
			Street string `json:"street" validate:"required,max=50"`
			Suburb string `json:"suburb" validate:"required,max=50"`
			CityID uuid.UUID `json:"cityId" validate:"required"`
		}{}

		context.BindJSON(&receivedData)

		var validate *validator.Validate = validator.New()		
		validationErrors := validate.Struct(receivedData)		

		if validationErrors != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		addressEntity := businessEntities.Address{
			ID: uuid.NewV4(),
			IndoorNumber: receivedData.IndoorNumber,
			OutdoorNumber: receivedData.OutdoorNumber,
			Street: receivedData.Street,
			Suburb: receivedData.Suburb,
			CityID: receivedData.CityID,
			ServiceRequesterID: serviceRequesterID,
		}		

		addressMgtService := addressManagementService.AddressManagementService{}
		databaseError := addressMgtService.Register(addressEntity)

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		response := struct {
			ID uuid.UUID `json:"id"`
			IndoorNumber string `json:"indoorNumber"`
			OutdoorNumber string `json:"outdoorNumber"`
			Street string `json:"street"`
			Suburb string `json:"suburb"`
			CityID uuid.UUID `json:"cityId"`
		}{
			ID: addressEntity.ID,
			IndoorNumber: addressEntity.IndoorNumber,
			OutdoorNumber: addressEntity.OutdoorNumber,
			Street: addressEntity.Street,
			Suburb: addressEntity.Suburb,
			CityID: addressEntity.CityID,
		}

		context.JSON(http.StatusCreated, response)
	}
}

func FindAllAddressesOfServiceRequester() gin.HandlerFunc {
	return func(context *gin.Context){
		serviceRequesterID, parsingError := uuid.FromString(context.Param("requesterId"))

		if parsingError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		// var addresses []businessEntities.Address
		// addressRepository := addressRepository.AddressRepository{}
		// databaseError := addressRepository.FindMatches(&addresses, "service_requester_id = ?", serviceRequesterID)

		// if databaseError != nil {
		// 	context.AbortWithStatus(http.StatusConflict)
		// 	return
		// }

		
		addressMgtService := addressManagementService.AddressManagementService{}
		addresses, databaseError := addressMgtService.FindAll(serviceRequesterID)

		if databaseError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}


		response := []struct {
			ID uuid.UUID `json:"id"`
			IndoorNumber string `json:"indoorNumber"`
			OutdoorNumber string `json:"outdoorNumber"`
			Street string `json:"street"`
			Suburb string `json:"suburb"`
			City struct{
				ID uuid.UUID `json:"id"`
				Name string `json:"name"`
			}
		}{}

		for _, address := range addresses {
			response = append(response, struct{ID uuid.UUID "json:\"id\""; IndoorNumber string "json:\"indoorNumber\""; OutdoorNumber string "json:\"outdoorNumber\""; Street string "json:\"street\""; Suburb string "json:\"suburb\""; City struct{ID uuid.UUID "json:\"id\""; Name string "json:\"name\""}}{
				ID: address.ID,
				IndoorNumber: address.IndoorNumber,
				OutdoorNumber: address.OutdoorNumber,
				Street: address.Street,
				Suburb: address.Suburb,
				City: struct{ID uuid.UUID "json:\"id\""; Name string "json:\"name\""}{
					ID: address.City.ID,
					Name: address.City.Name,
				},
			})
		}
		context.JSON(http.StatusOK, response)
	}
}