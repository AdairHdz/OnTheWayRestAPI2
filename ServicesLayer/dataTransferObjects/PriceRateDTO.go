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
	Price float32 `json:"price" validate:"min=1,max=100"`
	KindOfService uint8 `json:"kindOfService" validate:"min=0,max=4"`
	CityID uuid.UUID `json:"cityId"`
	WorkingDays []int `json:"workingDays" validate:"required,unique,min=1,max=7"`
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