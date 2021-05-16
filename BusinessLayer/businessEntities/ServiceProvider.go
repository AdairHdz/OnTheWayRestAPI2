package businessEntities

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	gorm.Model
	ID uuid.UUID
	User User
	UserID uuid.UUID `gorm:"size:191"`
	AverageScore float32
	Reviews []Review
	PriceRates []PriceRate
}