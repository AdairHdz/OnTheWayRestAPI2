package dataTransferObjects

import (
	"time"

	uuid "github.com/satori/go.uuid")


type ReceivedServiceRequestDTO struct {
	Cost              float32 `json:"cost"`
	DeliveryAddressID uuid.UUID `json:"deliveryAddressId"`
	Description string `json:"description"`
	KindOfService uint8 `json:"kindOfService"`
	ServiceProviderID uuid.UUID `json:"serviceProviderId"`
	ServiceRequesterID uuid.UUID `json:"serviceRequesterId"`
}

type ResponseServiceRequestDTO struct {
	ID uuid.UUID `json:"id"`
	Date time.Time `json:"date"`
	Status uint8 `json:"status"`
	Cost              float32 `json:"cost"`
	DeliveryAddressID uuid.UUID `json:"deliveryAddressId"`
	Description string `json:"description"`
	KindOfService uint8 `json:"kindOfService"`
	ServiceProviderID uuid.UUID `json:"serviceProviderId"`
	ServiceRequesterID uuid.UUID `json:"serviceRequesterId"`
}

type ResponseServiceRequestDTOWithDetails struct {
	ID uuid.UUID `json:"id"`
	Date time.Time `json:"date"`
	Status uint8 `json:"status"`
	Cost              float32 `json:"cost"`
	DeliveryAddress ResponseAddressWithCityDTO `json:"deliveryAddressId"`
	Description string `json:"description"`
	KindOfService uint8 `json:"kindOfService"`
	ServiceProvider ResponseUserDTOWithNamesOnly `json:"serviceProviderId"`
	ServiceRequester ResponseUserDTOWithNamesOnly `json:"serviceRequesterId"`
}