package businessEntities

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type State struct {
	gorm.Model
	ID uuid.UUID
	Name string
	Cities []City
}