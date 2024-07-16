package main

import (
	"ama/internal/data"
	"ama/internal/utils"
	"fmt"
	"time"

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
		statusText := utils.BoolToStatusDisplayText(employee.IsPresent)
		fmt.Printf("**** %s's attendance with ID '%s' has a status of '%s' on %s\n",
			employee.Name, employee.ID, statusText, utils.FormatDate(employee.CreatedAt))
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
		isPresent, err = utils.StatusInputTextToBool(presenceStatus)
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
		CreatedAt: time.Now(),
	}
	savedEmployee := app.models.Employees.Add(employee)
	statusText := utils.BoolToStatusDisplayText(savedEmployee.IsPresent)
	fmt.Printf("(%s's attendance is saved with ID '%s' & status '%s' on %s)\n",
		savedEmployee.Name, employee.ID, statusText, utils.FormatDate(employee.CreatedAt))
}

func (app *application) updateEmployeeStatus(parts []string) {
	if len(parts) != 3 {
		fmt.Println(`input should be \u [employee-id] [status (y/n)]`)
		return
	}

	employeeID := parts[1]
	presenceStatus := parts[2]

	isPresent, err := utils.StatusInputTextToBool(presenceStatus)
	if err != nil {
		fmt.Println(err)
		return
	}

	updatedEmployee, err := app.models.Employees.UpdatePresenceStatus(employeeID, isPresent)
	if err != nil {
		fmt.Println(err)
		return
	}

	statusText := utils.BoolToStatusDisplayText(updatedEmployee.IsPresent)
	fmt.Printf("(presence status of '%s' on %s is updated to '%s')\n",
		updatedEmployee.Name, utils.FormatDate(updatedEmployee.CreatedAt), statusText)
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

	fmt.Printf("(attendance data with ID '%s' is deleted)\n", employeeID)
}
