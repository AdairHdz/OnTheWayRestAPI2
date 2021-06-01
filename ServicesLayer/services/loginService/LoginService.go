package loginService

import (
	"net/http"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/tokenGenerator"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)


type LoginService struct{}

func (LoginService) Login()  gin.HandlerFunc {	
	return func(context *gin.Context){
		receivedData := struct {
			EmailAddress string
			Password string
		}{}

		context.BindJSON(&receivedData)
					
		user := businessEntities.User{
			EmailAddress: receivedData.EmailAddress,
		}
	
		userTypeID, loginError := user.Login()			

		if loginError != nil {
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		if user.ID == uuid.Nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		passwordError := hashing.VerifyPassword(user.Password, receivedData.Password)
		if passwordError != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, tokenError := tokenGenerator.CreateToken(user.EmailAddress, int(user.UserType))

		if tokenError != nil {
			context.Status(http.StatusConflict)
			return
		}

		refreshToken, tokenError := tokenGenerator.CreateRefreshToken(user.EmailAddress, int(user.UserType))

		if tokenError != nil {
			context.Status(http.StatusConflict)
			return
		}		

		response := struct {
			ID uuid.UUID `json:"id"`
			Names string `json:"names"`
			LastName string `json:"lastName"`
			EmailAddress string `json:"emailAddress"`
			UserType uint8 `json:"userType"`
			Verified bool `json:"verified"`
			StateID uuid.UUID `json:"stateId"`
			Token string `json:"token"`
			RefreshToken string `json:"refreshToken"`
		}{
			ID: userTypeID,
			Names: user.Names,
			LastName: user.LastName,
			EmailAddress: user.EmailAddress,
			UserType: user.UserType,
			Verified: user.Verified,
			StateID: user.StateID,
			Token: token,
			RefreshToken: refreshToken,
		}

		context.JSON(http.StatusOK, response)
	}
}