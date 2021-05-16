package addressRepository

import "github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"

type AddressRepository struct{}

func (AddressRepository) FindMatches(target interface{}, query interface{}, args interface{}) error {
	DB := database.GetDatabase()
	result := DB.Preload("City").Where(query, args).Find(target)
	return result.Error
}