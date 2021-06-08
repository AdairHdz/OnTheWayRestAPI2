package businessEntities

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories/serviceRequesterRepository"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ServiceRequester struct {
	gorm.Model
	ID        uuid.UUID
	User      User
	UserID    uuid.UUID `gorm:"size:191"`
	Addresses []Address
}

func (serviceRequester ServiceRequester) Register() error {
	repository := repositories.Repository{}
	databaseError := repository.Create(&serviceRequester)
	return databaseError
}

func (serviceRequester *ServiceRequester) Find(serviceRequesterID uuid.UUID) error {
	repository := serviceRequesterRepository.ServiceRequesterRepository{}
	databaseError := repository.FindByID(&serviceRequester, serviceRequesterID)

	return databaseError

}

func (serviceRequester *ServiceRequester) Update() error {
	repository := serviceRequesterRepository.ServiceRequesterRepository{}
	databaseError := repository.Update(&serviceRequester.User)
	return databaseError

}

func (ServiceRequester *ServiceRequester) GetStatisticsReport(requestedServicesPerWeekdayqueryResult, kindOfServicesQueryResult interface{}, startingDate, endingDate string) error {
	repository := serviceRequesterRepository.ServiceRequesterRepository{}

	// statisticsReport := struct {
	// 	RequestedServicesPerWeekdayqueryResult []struct {
	// 		RequestedServices int
	// 		Weekday           int
	// 	}

	// 	KindOfServicesQueryResult []struct {
	// 		RequestedServices int
	// 		KindOfService     int
	// 	}
	// }{}

	databaseError := repository.GetStatisticsReport(requestedServicesPerWeekdayqueryResult, kindOfServicesQueryResult, ServiceRequester.ID.String(), startingDate, endingDate)
	return databaseError
}
