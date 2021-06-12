package dataTransferObjects

import (
	uuid "github.com/satori/go.uuid"
)

type ReceivedServiceRequestDTO struct {
	Cost               float32   `json:"cost" validate:"min=1,max=100"`
	DeliveryAddressID  uuid.UUID `json:"deliveryAddressId"`
	Description        string    `json:"description" validate:"max=250"`
	KindOfService      uint8     `json:"kindOfService" validate:"min=0,max=4"`
	ServiceProviderID  uuid.UUID `json:"serviceProviderId"`
	ServiceRequesterID uuid.UUID `json:"serviceRequesterId"`
}

type ResponseServiceRequestDTO struct {
	ID                 uuid.UUID `json:"id"`
	Date               string    `json:"date"`
	Status             uint8     `json:"status"`
	Cost               float32   `json:"cost"`
	DeliveryAddressID  uuid.UUID `json:"deliveryAddressId"`
	Description        string    `json:"description"`
	KindOfService      uint8     `json:"kindOfService"`
	ServiceProviderID  uuid.UUID `json:"serviceProviderId"`
	ServiceRequesterID uuid.UUID `json:"serviceRequesterId"`
}

type ResponseServiceRequestDTOWithDetails struct {
	ID               uuid.UUID                    `json:"id"`
	Date             string                       `json:"date"`
	Status           uint8                        `json:"status"`
	Cost             float32                      `json:"cost"`
	DeliveryAddress  ResponseAddressWithCityDTO   `json:"deliveryAddress"`
	Description      string                       `json:"description"`
	KindOfService    uint8                        `json:"kindOfService"`
	ServiceProvider  ResponseUserDTOWithNamesOnly `json:"serviceProvider"`
	ServiceRequester ResponseUserDTOWithNamesOnly `json:"serviceRequester"`
}
