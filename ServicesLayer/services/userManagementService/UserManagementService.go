package userManagementService

import (
	"net/http"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type UserManagementService struct{}

func (UserManagementService) VerifyAccount() gin.HandlerFunc {
	return func(context *gin.Context) {
		userID, parsingError := uuid.FromString(context.Param("userId"))
		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The user ID you provided is not valid.")
			return
		}

		receivedData := struct {
			VerificationCode string `json:"verificationCode" validate:"len=8"`			
		}{}
		
		context.BindJSON(&receivedData)

		validator := validators.GetValidator()		
		validationError := validator.Struct(receivedData)

		if validationError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The verification code you sent has a non-valid format.")
			return
		}

		userWithAccountToBeVerified := businessEntities.User{			
			VerificationCode: "",
			Verified: true,
		}

		activationError := userWithAccountToBeVerified.VerifyAccount(userID.String(), receivedData.VerificationCode)
		
		if activationError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to activate your account.")
			return
		}

		context.Status(http.StatusOK)

	}
}