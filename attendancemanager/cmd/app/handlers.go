package main

import (
	"attendancemanager/internal/data"
	"attendancemanager/internal/utils"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (app *application) getAllEmployees() {
	for _, employee := range app.models.Employees.GetAll() {
		statusText := utils.GenerateStatusText(employee.IsPresent)

		fmt.Printf("** %s with ID '%s' is %s\n", employee.Name, employee.ID, statusText)
	}
}

func (app *application) addEmployee() {
	var name string
	var isPresent bool

	for {
		name = utils.GetInput(app.scanner, ">> Enter employee name: ")
		if name == "" {
			fmt.Println("name is required")
			continue
		}
		break
	}

	for {
		presenceStatus := utils.GetInput(app.scanner, ">> Is the employee present? [y/n] (default: n) ")
		if strings.ToLower(presenceStatus) == "y" {
			isPresent = true
		} else if strings.ToLower(presenceStatus) == "n" || presenceStatus == "" {
			isPresent = false
		} else {
			fmt.Println("presence status should be 'y', 'n', or empty")
			continue
		}
		break
	}

	employee := data.Employee{
		ID:        uuid.NewString(),
		Name:      name,
		IsPresent: isPresent,
	}
	savedEmployee := app.models.Employees.Add(employee)
	statusText := utils.GenerateStatusText(savedEmployee.IsPresent)
	fmt.Printf("'%s' with ID '%s' & presence status '%s' is saved\n", savedEmployee.Name, employee.ID, statusText)
}
