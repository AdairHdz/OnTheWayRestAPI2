package serviceRequestManagementService

import "github.com/gin-gonic/gin"

type IServiceRequestManagement interface {
	Register() gin.HandlerFunc
	FindByID() gin.HandlerFunc
	FindByDate(userType int) gin.HandlerFunc
	Update() gin.HandlerFunc
}
