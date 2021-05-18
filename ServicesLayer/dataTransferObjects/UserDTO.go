package dataTransferObjects

import (	
	uuid "github.com/satori/go.uuid"
)

type ReceivedUserDTO struct {
	Names        string    `json:"names" validate:"required,min=1,max=50,lettersAndSpaces"`
	LastName     string    `json:"lastName" validate:"required,min=1,max=50,lettersAndSpaces"`
	EmailAddress string    `json:"emailAddress" validate:"required,email,max=254"`	
	Password     string    `json:"password" validate:"required,max=80"`
	UserType uint8 `json:"userType" validate:"min=0,max=1"`
	StateID      uuid.UUID `json:"stateId" validate:"required"`
}

type ResponseUserDTO struct {
	ID uuid.UUID `json:"id"`
	Names        string    `json:"names"`
	LastName     string    `json:"lastName"`
	EmailAddress string    `json:"emailAddress"`
	UserType uint8 `json:"userType"`
	Verified bool `json:"verified"`
	StateID      uuid.UUID `json:"stateId"`
}

type ResponseUserDTOWithNamesOnly struct {
	ID uuid.UUID `json:"id"`
	Names string `json:"names"`
	LastName string `json:"lastName"`
}