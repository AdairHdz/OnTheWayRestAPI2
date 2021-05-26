package customErrors

type RecordNotFoundError struct { }

func (recordNotFoundError RecordNotFoundError) Error() string {
	return "Record not found"
}