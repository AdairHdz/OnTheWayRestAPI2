package serviceRequesterManagementService

import (
	"github.com/gin-gonic/gin"	
)


type IServiceRequesterManagementService interface {
	Register() gin.HandlerFunc
	Find() gin.HandlerFunc
	Update() gin.HandlerFunc
}
