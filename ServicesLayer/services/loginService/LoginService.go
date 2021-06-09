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

func (LoginService) Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		receivedData := struct {
			EmailAddress string
			Password     string
		}{}

		context.BindJSON(&receivedData)

		user := businessEntities.User{
			EmailAddress: receivedData.EmailAddress,
		}

		userTypeID, loginError := user.Login()

		if loginError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to verify your credentials. Please try again later.")
			return
		}

		if user.ID == uuid.Nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid credentials.")
			return
		}

		passwordError := hashing.VerifyPassword(user.Password, receivedData.Password)
		if passwordError != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid credentials.")
			return
		}

		token, tokenError := tokenGenerator.CreateToken(user.EmailAddress, int(user.UserType))

		if tokenError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to log you in. Please try again later.")
			return
		}

		refreshToken, tokenError := tokenGenerator.CreateRefreshToken(user.EmailAddress, int(user.UserType))

		if tokenError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to log you in. Please try again later.")
			return
		}

		response := struct {
			ID           uuid.UUID `json:"id"`
			UserID       uuid.UUID `json:"userId"`
			Names        string    `json:"names"`
			LastName     string    `json:"lastName"`
			EmailAddress string    `json:"emailAddress"`
			UserType     uint8     `json:"userType"`
			Verified     bool      `json:"verified"`
			StateID      uuid.UUID `json:"stateId"`
			Token        string    `json:"token"`
			RefreshToken string    `json:"refreshToken"`
		}{
			ID:           userTypeID,
			UserID:       user.ID,
			Names:        user.Names,
			LastName:     user.LastName,
			EmailAddress: user.EmailAddress,
			UserType:     user.UserType,
			Verified:     user.Verified,
			StateID:      user.StateID,
			Token:        token,
			RefreshToken: refreshToken,
		}

		context.JSON(http.StatusOK, response)
	}
}
