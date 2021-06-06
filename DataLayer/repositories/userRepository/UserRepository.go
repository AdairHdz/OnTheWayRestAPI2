package userRepository

import (
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"
	"github.com/AdairHdz/OnTheWayRestAPI/helpers/customErrors"
)


type UserRepository struct{}

func (UserRepository) VerifyAccount(userID, verification string, entity interface{}) (error) {
	DB := database.GetDatabase()
	result := DB.Where("id = ? AND verification_code = ?", userID, verification).Updates(entity)
	if result.RowsAffected == 0 {
		return customErrors.RecordNotFoundError{}
	}
	return result.Error
}