package businessEntities

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type City struct {
	gorm.Model
	ID uuid.UUID
	Name string
	StateID uuid.UUID `gorm:"size:191"`
}