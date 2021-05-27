package businessEntities

import (
	"errors"
	"time"

	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/repositories/priceRateRepository"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/customErrors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)


type PriceRate struct {
	gorm.Model
	ID uuid.UUID
	StartingHour string
	EndingHour string
	Price float32
	WorkingDays []WorkingDay `gorm:"many2many:pricerate_workingday;constraint:OnDelete:CASCADE;"`
	ServiceProviderID uuid.UUID `gorm:"size:191"`
	CityID uuid.UUID `gorm:"size:191"`
	City City
	KindOfService uint8
}

func (priceRate PriceRate) Register() error {
	repository := repositories.Repository{}
	priceRates, searchError := priceRate.Find(priceRate.ServiceProviderID)

	if searchError != nil {
		return searchError
	}

	
	for _, priceRateItem := range priceRates {
		startingHourOfExistentPriceRate, err := time.Parse(time.Kitchen, priceRateItem.StartingHour)
		if err != nil{
			return err
		}

		endingHourOfExistentPriceRate, err := time.Parse(time.Kitchen, priceRateItem.EndingHour)
		if err != nil{
			return err
		}

		shceduleCollides, err := priceRate.priceRateShceduleCollides(startingHourOfExistentPriceRate, endingHourOfExistentPriceRate)		
		if priceRateItem.appliesInTheSameCity(priceRate.CityID) &&
		priceRateItem.appliesToTheSameKindOfService(priceRate.KindOfService) &&
		priceRateItem.hasAtLeastOneOfTheInputWorkingDays(priceRate.WorkingDays) && shceduleCollides{			
			return errors.New("Schedule collision")
		}
	}
	
	databaseError := repository.Create(&priceRate)
	return databaseError
}

func (priceRate PriceRate) priceRateShceduleCollides(startingHourOfExistingPriceRate, endingHourOfExistingPriceRate time.Time) (bool, error) {

	startingHourOfNewPriceRate, err := time.Parse(time.Kitchen, priceRate.StartingHour)
	if err != nil{
		return true, err
	}

	endingHourOfNewPriceRate, err := time.Parse(time.Kitchen, priceRate.EndingHour)
	if err != nil {
		return true, err
	}

	if startingHourOfNewPriceRate.Sub(endingHourOfNewPriceRate) == 0{
		return true, nil
	}

	if startingHourOfNewPriceRate.Sub(startingHourOfExistingPriceRate) == 0 || startingHourOfNewPriceRate.Sub(endingHourOfExistingPriceRate) == 0 {
		return true, nil
	}
	
	if endingHourOfNewPriceRate.Sub(startingHourOfExistingPriceRate) == 0 || endingHourOfNewPriceRate.Sub(endingHourOfExistingPriceRate) == 0{
		return true, nil
	}
	
	if startingHourOfNewPriceRate.After(startingHourOfExistingPriceRate) && startingHourOfNewPriceRate.Before(endingHourOfExistingPriceRate) {
		return true, nil
	}
	
	if endingHourOfNewPriceRate.After(startingHourOfExistingPriceRate) && endingHourOfNewPriceRate.Before(endingHourOfExistingPriceRate) {
		return true, nil
	}
	
	if startingHourOfNewPriceRate.Before(startingHourOfExistingPriceRate) && endingHourOfNewPriceRate.After(endingHourOfExistingPriceRate) {
		return true, nil
	}
	
	return false, nil
}

func (priceRate PriceRate) appliesInTheSameCity(cityID uuid.UUID) bool{
	return priceRate.CityID == cityID
}

func (priceRate PriceRate) appliesToTheSameKindOfService(kindOfService uint8) bool{
	return priceRate.KindOfService == kindOfService
}

func (priceRate PriceRate) hasAtLeastOneOfTheInputWorkingDays(workingDays []WorkingDay) bool{
	for _, workingDayElement := range priceRate.WorkingDays {
		for _, inputWorkingDayElement := range workingDays {
			if workingDayElement.ID == inputWorkingDayElement.ID {
				return true
			}
		}
	}
	return false
}

func (priceRate PriceRate) Find(serviceProviderID uuid.UUID) ([]PriceRate, error) {
	repository := priceRateRepository.PriceRateRepository{}
	var priceRates []PriceRate
	databaseError :=  repository.FindMatches(&priceRates, "service_provider_id = ?", serviceProviderID)
	return priceRates, databaseError	
}

func (priceRate *PriceRate) Delete(serviceProviderID uuid.UUID) error {
	repository := repositories.Repository{}		
	rowsAffected, databaseError := repository.Delete(&priceRate, "service_provider_id = ?", serviceProviderID)		
	if rowsAffected == 0 {
		return customErrors.RecordNotFoundError{}
	}
	return databaseError
}