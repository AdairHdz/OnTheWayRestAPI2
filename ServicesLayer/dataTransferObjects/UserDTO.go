package dataTransferObjects

import (	
	uuid "github.com/satori/go.uuid"
)

type UserDTO struct {
	ID uuid.UUID `json:"id"`
	Names        string    `json:"names" validate:"required,min=1,max=50,lettersAndSpaces"`
	LastName     string    `json:"lastName" validate:"required,min=1,max=50, lettersAndSpaces"`
	EmailAddress string    `json:"emailAddress" validate:"required,email,max=254"`
	UserType uint8 `json:"userType"`
	Verified bool `json:"verified"`
	Password     string    `json:"password" validate:"required,max=80"`
	StateID      uuid.UUID `json:"stateId" validate:"required"`
}