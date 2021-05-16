package businessEntities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type PriceRateWorkingDay struct {
	PriceRateID uuid.UUID `gorm:"primaryKey;size:191"`
	WorkingDayID uint8 `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
