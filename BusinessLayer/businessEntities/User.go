package businessEntities

import (	
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const (
	ServiceProviderType = iota
	ServiceRequesterType = iota
)

type User struct {
	gorm.Model
	ID uuid.UUID
	Names        string 
	LastName     string 
	EmailAddress string `gorm:"unique"`
	Password     string 
	Verified     bool   
	UserType     uint8
	VerificationCode string
	StateID      uuid.UUID `gorm:"size:191"`
	State        State
}

func (user *User) Login() (uuid.UUID, error) {
	repository := repositories.Repository{}
	databaseError := repository.FindMatches(&user, "email_address = ?", user.EmailAddress)
	var userTypeID uuid.UUID	
	if databaseError == nil{
		if user.UserType == ServiceProviderType{
			var serviceProvider ServiceProvider
			repository.FindMatches(&serviceProvider, "user_id = ?", user.ID)			
			userTypeID = serviceProvider.ID
		}else{
			var serviceRequester ServiceRequester
			repository.FindMatches(&serviceRequester, "user_id = ?", user.ID)
			userTypeID = serviceRequester.ID			
		}
	}	
	return userTypeID, databaseError
}
