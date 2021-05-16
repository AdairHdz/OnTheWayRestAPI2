package mappers

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/dataTransferObjects"
	uuid "github.com/satori/go.uuid"
)


func CreateAddressDTOAsResponse(address businessEntities.Address) dataTransferObjects.ResponseAddressDTO {
	addressDTO := dataTransferObjects.ResponseAddressDTO {
		ID: address.ID,
		IndoorNumber: address.IndoorNumber,
		OutdoorNumber: address.OutdoorNumber,
		Street: address.Street,
		Suburb: address.Suburb,
		CityID: address.CityID,
	}

	return addressDTO
}

func CreateAddressDTOWithCityAsResponse(address businessEntities.Address) dataTransferObjects.ResponseAddressWithCityDTO {
	addressDTO := dataTransferObjects.ResponseAddressWithCityDTO{
		ID: address.ID,
		IndoorNumber: address.IndoorNumber,
		OutdoorNumber: address.OutdoorNumber,
		Street: address.Street,
		Suburb: address.Suburb,
		City: dataTransferObjects.CityDTO{
			ID: address.City.ID,
			Name: address.City.Name,
		},
	}

	return addressDTO
}

func CreateAddressEntity(addressDTO dataTransferObjects.ReceivedAddressDTO, serviceRequesterID uuid.UUID) businessEntities.Address{
	addressEntity := businessEntities.Address{
		ID: uuid.NewV4(),
		IndoorNumber: addressDTO.IndoorNumber,
		OutdoorNumber: addressDTO.OutdoorNumber,
		Street: addressDTO.Street,
		Suburb: addressDTO.Suburb,
		CityID: addressDTO.CityID,
		ServiceRequesterID: serviceRequesterID,
	}
	return addressEntity
}