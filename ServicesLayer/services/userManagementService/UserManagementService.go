package userManagementService

import (
	"net/http"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/validators"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/codeGenerator"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/mailerService"
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

func (UserManagementService) GetNewVerificationCode() gin.HandlerFunc {
	return func(context *gin.Context) {
		userID, parsingError := uuid.FromString(context.Param("userId"))
		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The user ID you provided is not valid.")
			return
		}

		receivedData := struct {
			EmailAddress string `json:"emailAddress" validate:"required,email,max=254"`
		}{}

		context.BindJSON(&receivedData)

		validator := validators.GetValidator()
		validationError := validator.Struct(receivedData)

		if validationError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The email address you provided is not valid.")
			return
		}

		newVerificationCode := codeGenerator.GenerateCode()

		verificationMailer := mailerService.MailerGRPCService{}
		verificationMailer.SendEmail(receivedData.EmailAddress, newVerificationCode)

		user := businessEntities.User{			
			VerificationCode: newVerificationCode,
		}

		refreshingCodeError := user.RefreshVerificationCode(userID.String())

		if refreshingCodeError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to get a new verification code.")
			return
		}
		
		context.Status(http.StatusOK)
	}
}