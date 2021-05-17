package dataTransferObjects

import (
	"time"
	uuid "github.com/satori/go.uuid")


type ResponsePriceRateDTO struct {
	ID uuid.UUID `json:"id"`
	StartingHour time.Time `json:"startingHour"`
	EndingHour time.Time `json:"endingHour"`
	Price float32 `json:"price"`
	KindOfService uint8 `json:"kindOfService"`
	City CityDTO `json:"city"`
	WorkingDays []ResponseWorkingDayDTO `json:"workingDays"`
}