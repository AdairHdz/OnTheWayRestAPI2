package serviceProviderManagementService

import "github.com/gin-gonic/gin"

type IServiceProviderManagementService interface {
	Find() gin.HandlerFunc
	FindMatches() gin.HandlerFunc
	Update() gin.HandlerFunc
	UpdateServiceProviderImage() gin.HandlerFunc
	GetStatistics() gin.HandlerFunc
}
