package data

import "errors"

var (
	ErrEmployeeNotFound = errors.New("employee not found")
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
