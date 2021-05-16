package serviceRequesterRepository

import (
"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"
)


type ServiceRequesterRepository struct{}

func (ServiceRequesterRepository) FindByID(serviceRequester interface{}, ID interface{}) (error) {
	DB := database.GetDatabase()
	result := DB.Preload("User").Preload("User.State").First(serviceRequester, ID)	
	return result.Error
}

func (ServiceRequesterRepository) Update(serviceRequester interface{}) (error) {
	DB := database.GetDatabase()
	result := DB.Updates(serviceRequester)
	return result.Error	
}