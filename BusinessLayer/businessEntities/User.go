package businessEntities

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories/userRepository"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const (
	ServiceProviderType  = iota
	ServiceRequesterType = iota
)

type User struct {
	gorm.Model
	ID               uuid.UUID
	Names            string
	LastName         string
	EmailAddress     string `gorm:"unique"`
	Password         string
	Verified         bool
	UserType         uint8
	VerificationCode string
	RecoveryCode     string
	StateID          uuid.UUID `gorm:"size:191"`
	State            State
}

func (user *User) Login() (uuid.UUID, error) {
	repository := repositories.Repository{}
	databaseError := repository.FindMatches(&user, "email_address = ?", user.EmailAddress)
	var userTypeID uuid.UUID
	if databaseError == nil {
		if user.UserType == ServiceProviderType {
			var serviceProvider ServiceProvider
			repository.FindMatches(&serviceProvider, "user_id = ?", user.ID)
			userTypeID = serviceProvider.ID
		} else {
			var serviceRequester ServiceRequester
			repository.FindMatches(&serviceRequester, "user_id = ?", user.ID)
			userTypeID = serviceRequester.ID
		}
	}
	return userTypeID, databaseError
}

func (user *User) VerifyAccount(userID, activationCode string) error {
	repository := userRepository.UserRepository{}
	databaseError := repository.VerifyAccount(userID, activationCode, user)
	return databaseError
}

func (user *User) RefreshVerificationCode(userID string) error {
	repository := userRepository.UserRepository{}
	databaseError := repository.RefreshVerificationCode(userID, user)
	return databaseError
}

func (user *User) RecoverPassword(userID string, recoveryCode string) error {
	repository := userRepository.UserRepository{}
	databaseError := repository.RecoverPassword(userID, recoveryCode, user)
	return databaseError
}

func (user *User) RefreshRecoveryCode(emailAddress string) error {
	repository := userRepository.UserRepository{}
	databaseError := repository.RefreshRecoveryCode(emailAddress, user)
	return databaseError
}
