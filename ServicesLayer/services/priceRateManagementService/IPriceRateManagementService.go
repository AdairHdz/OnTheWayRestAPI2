package priceRateManagementService

import "github.com/gin-gonic/gin"

type IPriceRateManagementService interface {
	Register() gin.HandlerFunc
	FindAll() gin.HandlerFunc
	Delete() gin.HandlerFunc
}