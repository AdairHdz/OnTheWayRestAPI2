package reviewRepository

import "github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"

type ReviewRepository struct{}

func (ReviewRepository) FindMatches(target interface{}, query interface{}, args interface{}) error {
	DB := database.GetDatabase()
	result := DB.Preload("ServiceRequester").Preload("ServiceRequester.User").Preload("Evidence").Where(query, args).Find(target)
	return result.Error
}