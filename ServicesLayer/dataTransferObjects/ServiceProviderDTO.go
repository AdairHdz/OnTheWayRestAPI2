package dataTransferObjects

import uuid "github.com/satori/go.uuid"

type ResponseServiceProviderDTO struct {
	ID uuid.UUID `json:"id"`
	Names string
	LastName string
	EmailAddress string
	AverageScore uint8
	PriceRates []ResponsePriceRateDTOWithCity
}

type ResponseServiceProviderOverviewDTO struct {
	ID uuid.UUID `json:"id"`
	Names string
	LastName string	
	AverageScore uint8
	PriceRate float32
}