package serviceProviderManagementService

import (
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ServiceProviderManagementService struct { }

func (ServiceProviderManagementService) Register(serviceProvider businessEntities.ServiceProvider) error {
	registryError := serviceProvider.Register()
	return registryError
}

func (ServiceProviderManagementService) Find(serviceProviderID uuid.UUID) (businessEntities.ServiceProvider, error) {
	var serviceProvider businessEntities.ServiceProvider
	searchError := serviceProvider.Find(serviceProviderID)
	return serviceProvider, searchError	
}

func (ServiceProviderManagementService) FindMatches() gin.HandlerFunc {
	return func(context *gin.Context){
		
	}
}

func (ServiceProviderManagementService) Update() gin.HandlerFunc {
	return func(context *gin.Context){
		
	}
}
