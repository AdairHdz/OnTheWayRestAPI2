package dataTransferObjects

import uuid "github.com/satori/go.uuid"

type CityDTO struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}