package businessEntities

import (
	"encoding/json"

	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	ID      uuid.UUID
	Name    string
	StateID uuid.UUID `gorm:"size:191"`
}

func (City) FindAll(stateID uuid.UUID) ([]City, error) {
	var cities []City
	repository := repositories.Repository{}
	databaseError := repository.FindMatches(&cities, "state_id = ?", stateID)
	return cities, databaseError
}

func (city City) MarshalBinary() (data []byte, err error) {
	return json.Marshal(city)
}

func (city *City) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &city)
}
