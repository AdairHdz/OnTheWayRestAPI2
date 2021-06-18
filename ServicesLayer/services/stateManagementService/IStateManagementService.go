package stateManagementService

import "github.com/gin-gonic/gin"

type IStateManagementService interface {
	FindAll() gin.HandlerFunc
	FindAllCitiesOfState() gin.HandlerFunc
}
