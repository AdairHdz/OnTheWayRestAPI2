package serviceProviderManagementService

import (	
	"github.com/gin-gonic/gin"	
)


type ServiceProviderManagementService struct{}

func (ServiceProviderManagementService) Register() gin.HandlerFunc {
	return func(context *gin.Context) {
		
	}
}

func (ServiceProviderManagementService) Find() gin.HandlerFunc {
	return func(context *gin.Context){
		
	}
}

func (ServiceProviderManagementService) FindMatches() gin.HandlerFunc {
	return func(context *gin.Context){
		
	}
}

func (ServiceProviderManagementService) Update() gin.HandlerFunc {
	return func(context *gin.Context){
		
	}
}
