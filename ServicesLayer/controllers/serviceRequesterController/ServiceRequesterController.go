package serviceRequesterController

import (	
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/services/serviceRequesterManagementService"	
	"github.com/gin-gonic/gin"
	
)

var (
	serviceRequesterMgtService = serviceRequesterManagementService.ServiceRequesterManagementService{}
)

func FindServiceRequester() gin.HandlerFunc{
	return func(context *gin.Context){
		

	}
}

// func UpdateServiceRequester() gin.HandlerFunc{

// }