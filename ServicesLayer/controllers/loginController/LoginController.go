package loginController

import (
	"net/http"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/loginService"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/tokenGenerator"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)


func Login() gin.HandlerFunc {
	return func(context *gin.Context){
		receivedData := struct {
			EmailAddress string
			Password string
		}{}

		context.BindJSON(&receivedData)
					
		login := loginService.LoginService{}
		user, loginError := login.Login(receivedData.EmailAddress)			

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
			context.AbortWithStatus(http.StatusConflict)
			return
		}

		token, tokenError := tokenGenerator.CreateToken(user.EmailAddress)

		if tokenError != nil {
			context.Status(http.StatusConflict)
			return
		}

		response := struct {
			ID uuid.UUID
			Names string
			LastName string
			EmailAddress string
			UserType uint8
			Verified bool
			StateID uuid.UUID
			Token string
		}{
			ID: user.ID,
			Names: user.Names,
			LastName: user.LastName,
			EmailAddress: user.EmailAddress,
			UserType: user.UserType,
			Verified: user.Verified,
			StateID: user.StateID,
			Token: token,
		}

		context.JSON(http.StatusOK, response)
	}
}