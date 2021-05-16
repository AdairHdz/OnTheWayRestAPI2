package loginService

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities")


type LoginService struct{}

func (LoginService) Login(emailAddress string) (businessEntities.User, error) {	
	user := businessEntities.User{
		EmailAddress: emailAddress,
	}

	loginError := user.Login()
	return user, loginError
}