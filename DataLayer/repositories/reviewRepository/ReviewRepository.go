package reviewRepository

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/paginator"
)

type ReviewRepository struct{}

func (ReviewRepository) FindMatches(page, pagesize int, rowCount *int64, target interface{}, query interface{}, args interface{}) error {
	DB := database.GetDatabase()
	result := DB.Scopes(paginator.Paginate(page, pagesize)).Preload("ServiceRequester").Preload("ServiceRequester.User").Preload("Evidence").Where(query, args).Find(target)
	var targetCount interface{}
	DB.Table("reviews").Where(query, args).Scan(targetCount).Count(rowCount)
	return result.Error
}
