package repositories

import "github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"

type Repository struct { }

func (Repository) Create(entity interface{}) error {
	DB := database.GetDatabase()
	result := DB.Create(entity)
	return result.Error
}

func (Repository) FindByID(entity interface{}, ID interface{}) (interface{}, error) {		
	DB := database.GetDatabase()
	result := DB.First(&entity, ID)
	return entity, result.Error
}

func (Repository) FindMatches(target interface{}, query interface{}, args interface{}) error {
	DB := database.GetDatabase()
	result := DB.Where(query, args).Find(target)
	return result.Error
}

func (Repository) Update(entity interface{}) error {
	DB := database.GetDatabase()
	result := DB.Updates(entity)
	return result.Error
}

func (Repository) Delete(entity, query, args interface{}) error {
	DB := database.GetDatabase()
	result := DB.Where(query, args).Delete(entity)
	return result.Error
}
