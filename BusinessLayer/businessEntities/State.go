package businessEntities

import (
	"encoding/json"

	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	ID     uuid.UUID
	Name   string
	Cities []City
}

func (State) FindAll() ([]State, error) {
	var states []State
	repository := repositories.Repository{}
	databaseError := repository.FindAll(&states)
	return states, databaseError
}

func (state State) MarshalBinary() (data []byte, err error) {
	return json.Marshal(state)
}

func (state *State) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &state)
}
