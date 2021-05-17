package businessEntities

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type ReviewEvidence struct {
	gorm.Model
	ID uuid.UUID
	Name string
	ReviewID uuid.UUID `gorm:"size:191"`	
}