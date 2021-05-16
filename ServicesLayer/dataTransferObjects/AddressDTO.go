package dataTransferObjects

import uuid "github.com/satori/go.uuid"

type ReceivedAddressDTO struct {
	IndoorNumber  string    `json:"indoorNumber" validate:"max=8"`
	OutdoorNumber string    `json:"outdoorNumber" validate:"required,max=8"`
	Street        string    `json:"street" validate:"required,max=50"`
	Suburb        string    `json:"suburb" validate:"required,max=50"`
	CityID        uuid.UUID `json:"cityId" validate:"required"`
}

type ResponseAddressDTO struct {
	ID uuid.UUID `json:"id"`
	IndoorNumber string `json:"indoorNumber"`
	OutdoorNumber string `json:"outdoorNumber"`
	Street string `json:"street"`
	Suburb string `json:"suburb"`
	CityID uuid.UUID `json:"cityId"`
}

type ResponseAddressWithCityDTO struct {
	ID uuid.UUID `json:"id"`
	IndoorNumber string `json:"indoorNumber"`
	OutdoorNumber string `json:"outdoorNumber"`
	Street string `json:"street"`
	Suburb string `json:"suburb"`
	City CityDTO
}