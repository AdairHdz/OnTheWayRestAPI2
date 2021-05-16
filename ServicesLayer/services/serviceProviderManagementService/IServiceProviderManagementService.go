package serviceProviderManagementService

import "github.com/gin-gonic/gin"

type IServiceProviderManagementService interface {
	Register() gin.HandlerFunc
	Find() gin.HandlerFunc
	FindMatches() gin.HandlerFunc
	Update() gin.HandlerFunc
}