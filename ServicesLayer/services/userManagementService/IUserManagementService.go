package userManagementService

import "github.com/gin-gonic/gin"

type IUserManagementService interface {
	VerifyAccount() gin.HandlerFunc
	GetNewVerificationCode() gin.HandlerFunc
	RecoverPassword() gin.HandlerFunc
	SendRecoveryCode() gin.HandlerFunc
}
