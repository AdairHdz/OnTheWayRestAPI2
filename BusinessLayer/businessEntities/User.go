package businessEntities

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	ID uuid.UUID
	Names        string 
	LastName     string 
	EmailAddress string 
	Password     string 
	Verified     bool   
	UserType     uint8
	StateID      uuid.UUID `gorm:"size:191"`
	State        State
}

func (user *User) Login() error {
	repository := repositories.Repository{}
	databaseError := repository.FindMatches(&user, "email_address = ?", user.EmailAddress)
	return databaseError
}
