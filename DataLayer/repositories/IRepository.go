package repositories

type IRepository interface{
	Create(entity interface{}) error
	FindById(entity interface{}, ID interface{}) error
	Find(target interface{}, query interface{}, args interface{}) error
}

