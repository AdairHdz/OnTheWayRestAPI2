package dataTransferObjects

import uuid "github.com/satori/go.uuid"

type ResponseServiceProviderDTO struct {
	ID           uuid.UUID                      `json:"id"`
	Names        string                         `json:"names"`
	LastName     string                         `json:"lastName"`
	EmailAddress string                         `json:"emailAddress"`
	AverageScore uint8                          `json:"averageScore"`
	PriceRates   []ResponsePriceRateDTOWithCity `json:"priceRates"`
	ProfileImage string                         `json:"profileImage"`
}

type ResponseServiceProviderOverviewDTO struct {
	ID           uuid.UUID `json:"id"`
	Names        string    `json:"names"`
	LastName     string    `json:"lastName"`
	AverageScore uint8     `json:"averageScore"`
	PriceRate    float32   `json:"priceRate"`
}
