package priceRateRepository

import "github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"

type PriceRateRepository struct{}

func (PriceRateRepository) FindMatches(target interface{}, query interface{}, args interface{}) error {
	DB := database.GetDatabase()
	result := DB.Preload("City").Preload("WorkingDays").Where(query, args).Find(target)
	return result.Error
}
