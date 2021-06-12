package serviceRequesterRepository

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/customErrors"
)

type ServiceRequesterRepository struct{}

func (ServiceRequesterRepository) FindByID(serviceRequester interface{}, ID interface{}) error {
	DB := database.GetDatabase()
	result := DB.Preload("User").Preload("User.State").First(serviceRequester, ID)
	if result.RowsAffected == 0 {
		return customErrors.RecordNotFoundError{}
	}
	return result.Error
}

func (ServiceRequesterRepository) Update(serviceRequester interface{}) error {
	DB := database.GetDatabase()
	result := DB.Updates(serviceRequester)
	return result.Error
}

func (ServiceRequesterRepository) GetStatisticsReport(requestedServicesPerWeekdayqueryResult, kindOfServicesQueryResult interface{}, serviceRequesterID, startingDate, endingDate string) error {
	DB := database.GetDatabase()
	requestedServicesPerWeekdayDatabaseResult := DB.Raw("SELECT COUNT(`id`) AS 'requested_services', "+
		"WEEKDAY(DATE(`date`)) 'weekday' "+
		"FROM service_requests "+
		"WHERE service_requester_id = ? AND DATEDIFF(?, `date`) <= 0 "+
		"AND DATEDIFF(?, `date`) >= 0 "+
		"GROUP BY WEEKDAY(DATE(`date`));", serviceRequesterID, startingDate, endingDate).Scan(requestedServicesPerWeekdayqueryResult)

	if requestedServicesPerWeekdayDatabaseResult.Error != nil {
		return requestedServicesPerWeekdayDatabaseResult.Error
	}

	kindOfServicesDatabaseResult := DB.Raw("SELECT COUNT(`id`) AS 'requested_services', "+
		"`kind_of_service` AS 'kind_of_service' "+
		"FROM service_requests "+
		"WHERE service_requester_id = ? "+
		"AND DATEDIFF(?, `date`) <= 0 "+
		"AND DATEDIFF(?, `date`) >= 0 "+
		"GROUP BY(`kind_of_service`);", serviceRequesterID, startingDate, endingDate).Scan(kindOfServicesQueryResult)

	if kindOfServicesDatabaseResult.Error != nil {
		return kindOfServicesDatabaseResult.Error
	}
	return nil
}
