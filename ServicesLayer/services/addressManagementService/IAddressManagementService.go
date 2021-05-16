package addressManagementService

import "github.com/gin-gonic/gin"

type IAddressManagementService interface {
	Register() gin.HandlerFunc
	FindAll() gin.HandlerFunc
}