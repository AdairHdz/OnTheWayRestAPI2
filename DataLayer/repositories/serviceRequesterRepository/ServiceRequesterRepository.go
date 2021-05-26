package serviceRequesterRepository

import (
"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"
"github.com/AdairHdz/OnTheWayRestAPI/helpers/customErrors"
)


type ServiceRequesterRepository struct{}

func (ServiceRequesterRepository) FindByID(serviceRequester interface{}, ID interface{}) (error) {
	DB := database.GetDatabase()
	result := DB.Preload("User").Preload("User.State").First(serviceRequester, ID)
	if result.RowsAffected == 0 {
		return customErrors.RecordNotFoundError{}
	}
	return result.Error
}

func (ServiceRequesterRepository) Update(serviceRequester interface{}) (error) {
	DB := database.GetDatabase()
	result := DB.Updates(serviceRequester)
	return result.Error	
}