package dataTransferObjects

import uuid "github.com/satori/go.uuid"

type StateDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}