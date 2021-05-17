package dataTransferObjects

import uuid "github.com/satori/go.uuid"

type ResponseServiceProviderDTO struct {
	ID uuid.UUID `json:"id"`
	Names string
	LastName string
	EmailAddress string
	AverageScore uint8
	PriceRates []ResponsePriceRateDTO
}