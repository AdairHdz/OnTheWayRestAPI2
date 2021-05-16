package serviceRequestManagementService

import "github.com/gin-gonic/gin"

type IServiceRequestManagement interface {
	Register() gin.HandlerFunc
	Find() gin.HandlerFunc
	Update() gin.HandlerFunc
}