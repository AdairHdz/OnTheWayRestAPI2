package serviceRequesterManagementService

import (
	"github.com/gin-gonic/gin"
)

type IServiceRequesterManagementService interface {
	GetStatistics() gin.HandlerFunc
	Find() gin.HandlerFunc
	Update() gin.HandlerFunc
}
