package dataTransferObjects

import (	
	uuid "github.com/satori/go.uuid")


type ResponsePriceRateDTO struct {
	ID uuid.UUID `json:"id"`
	StartingHour string `json:"startingHour"`
	EndingHour string `json:"endingHour"`
	Price float32 `json:"price"`
	KindOfService uint8 `json:"kindOfService"`
	CityID uuid.UUID `json:"cityId"`
	WorkingDays []int `json:"workingDays"`
}

type ReceivedPriceRateDTO struct {	
	StartingHour string `json:"startingHour"`
	EndingHour string `json:"endingHour"`
	Price float32 `json:"price"`
	KindOfService uint8 `json:"kindOfService"`
	CityID uuid.UUID `json:"cityId"`
	WorkingDays []int `json:"workingDays"`
}

type ResponsePriceRateDTOWithCity struct {
	ID uuid.UUID `json:"id"`
	StartingHour string `json:"startingHour"`
	EndingHour string `json:"endingHour"`
	Price float32 `json:"price"`
	KindOfService uint8 `json:"kindOfService"`
	City CityDTO `json:"city"`
	WorkingDays []int `json:"workingDays"`
}