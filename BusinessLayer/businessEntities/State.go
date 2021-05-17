package businessEntities

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type State struct {
	gorm.Model
	ID uuid.UUID
	Name string
	Cities []City
}

func (State) FindAll() ([]State, error) {
	var states []State
	repository := repositories.Repository{}
	databaseError := repository.FindAll(&states)
	return states, databaseError
}