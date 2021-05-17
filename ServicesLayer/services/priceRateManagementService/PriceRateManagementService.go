package priceRateManagementService

import (
	"github.com/gin-gonic/gin"
)


type PriceRateManagementService struct{}

func (PriceRateManagementService) Register() gin.HandlerFunc {
	return func(context *gin.Context){
		// receivedData := struct {
		// 	Title string 
		// 	Details string
		// 	Score uint8
			
		// }{}
	}
}

func (PriceRateManagementService) FindAll() gin.HandlerFunc {
	return func(context *gin.Context){
		
	}
}

func (PriceRateManagementService) Delete() gin.HandlerFunc {
	return func(context *gin.Context){		
		
	}
}
