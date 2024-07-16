package main

import (
	"attendancemanager/internal/data"
	"attendancemanager/internal/utils"
	"fmt"

	"github.com/google/uuid"
)

func (app *application) getAllEmployees(parts []string) {
	if len(parts) != 1 {
		fmt.Println(`input should be \l`)
		return
	}

	employees := app.models.Employees.GetAll()
	total := len(employees)

	fmt.Printf("total: %d\n", total)

	if total == 0 {
		fmt.Println("there are no employees available")
		return
	}

	for _, employee := range employees {
		statusText := utils.BooleanToStatusText(employee.IsPresent)
		fmt.Printf("**** %s with ID '%s' is %s\n", employee.Name, employee.ID, statusText)
	}
}

func (app *application) addEmployee(parts []string) {
	if len(parts) != 1 {
		fmt.Println(`input should be \a`)
		return
	}

	var name string
	var isPresent bool
	var err error

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
		isPresent, err = utils.StatusTextToBoolean(presenceStatus)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	employee := &data.Employee{
		ID:        uuid.NewString(),
		Name:      name,
		IsPresent: isPresent,
	}
	savedEmployee := app.models.Employees.Add(employee)
	statusText := utils.BooleanToStatusText(savedEmployee.IsPresent)
	fmt.Printf("('%s' is saved with ID '%s' & presence status '%s')\n", savedEmployee.Name, employee.ID, statusText)
}

func (app *application) updateEmployeeStatus(parts []string) {
	if len(parts) != 3 {
		fmt.Println(`input should be \u [employee-id] [status (y/n)]`)
		return
	}

	employeeID := parts[1]
	presenceStatus := parts[2]

	isPresent, err := utils.StatusTextToBoolean(presenceStatus)
	if err != nil {
		fmt.Println(err)
		return
	}

	updatedEmployee, err := app.models.Employees.UpdatePresenceStatus(employeeID, isPresent)
	if err != nil {
		fmt.Println(err)
		return
	}

	statusText := utils.BooleanToStatusText(updatedEmployee.IsPresent)
	fmt.Printf("(presence status of '%s' is updated to '%s')\n", updatedEmployee.Name, statusText)
}

func (app *application) deleteEmployee(parts []string) {
	if len(parts) != 2 {
		fmt.Println(`input should be \d [employee-id]`)
		return
	}

	employeeID := parts[1]

	err := app.models.Employees.DeleteByID(employeeID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("(employee with ID '%s' is deleted)\n", employeeID)
}
