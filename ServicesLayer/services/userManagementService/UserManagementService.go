package userManagementService

import (
	"net/http"

	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/mailerService"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/codeGenerator"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/hashing"
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
			Verified:         true,
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

func (UserManagementService) RecoverPassword() gin.HandlerFunc {
	return func(context *gin.Context) {
		userID, parsingError := uuid.FromString(context.Param("userId"))
		if parsingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The user ID you provided is not valid.")
			return
		}

		receivedData := struct {
			RecoveryCode string `json:"recoveryCode"`
			NewPassword  string `json:"newPassword"`
		}{}

		bindingError := context.BindJSON(&receivedData)
		if bindingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
			return
		}

		if len(receivedData.NewPassword) < 8 {
			context.AbortWithStatusJSON(http.StatusBadRequest, "Your new password is insecure. Please try again using at least 8 digits.")
			return
		}

		hashedPassword, hashingError := hashing.GenerateHash(receivedData.NewPassword)
		if hashingError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to update your password.")
			return
		}
		user := businessEntities.User{
			Password: hashedPassword,
		}

		passwordUpdateError := user.RecoverPassword(userID.String(), receivedData.RecoveryCode)
		if passwordUpdateError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to update your password.")
			return
		}

		context.Status(http.StatusOK)
	}
}

func (UserManagementService) SendRecoveryCode() gin.HandlerFunc {
	return func(context *gin.Context) {
		receivedData := struct {
			EmailAddress string `json:"emailAddress"`
		}{}
		bindingError := context.BindJSON(&receivedData)
		if bindingError != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
			return
		}
		validator := validators.GetValidator()
		validationErrors := validator.Var(receivedData.EmailAddress, "email")
		if validationErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "The data you provided has a non-valid format.")
			return
		}

		recoveryCode := codeGenerator.GenerateCode()
		user := businessEntities.User{
			RecoveryCode: recoveryCode,
		}

		refreshingError := user.RefreshRecoveryCode(receivedData.EmailAddress)
		if refreshingError != nil {
			context.AbortWithStatusJSON(http.StatusConflict, "There was an error while trying to update your recovery code.")
			return
		}

		mailer := mailerService.MailerGRPCService{}
		mailer.SendEmail(receivedData.EmailAddress, recoveryCode)
		context.Status(http.StatusOK)
	}
}
