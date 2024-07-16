package data

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Employees EmployeeModel
}

func NewModels() *Models {
	return &Models{
		Employees: EmployeeModel{
			DB: make(map[string]Employee),
		},
	}
}
