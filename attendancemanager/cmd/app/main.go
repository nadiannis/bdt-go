package main

import (
	"attendancemanager/internal/data"
	"attendancemanager/internal/utils"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type application struct {
	scanner *bufio.Scanner
	models  data.Models
}

func main() {
	app := application{
		scanner: bufio.NewScanner(os.Stdin),
		models:  *data.NewModels(),
	}

	fmt.Println("======== ATTENDANCE MANAGER ========")
	fmt.Println("Choose options:")
	fmt.Println(`\l => Get list of employees`)
	fmt.Println(`\a => Add a new employee`)
	fmt.Println(`\u [employee-id] [presence-status] => Update employee presence status`)
	fmt.Println(`\d [employee-id] => Delete an employee`)
	fmt.Println(`\q => Quit`)

	employee1 := data.Employee{
		ID:        uuid.NewString(),
		Name:      "Nadia",
		IsPresent: true,
	}
	employee2 := data.Employee{
		ID:   uuid.NewString(),
		Name: "Neyla",
	}

	app.models.Employees.Add(employee1)
	app.models.Employees.Add(employee2)

	for {
		input := utils.GetInput(app.scanner, "> ")
		if input == "" {
			fmt.Println("input is required")
			continue
		}

		parts := strings.Fields(input)
		action := parts[0]

		switch action {
		case `\l`:
			app.getAllEmployees()
		case `\a`:
			app.addEmployee()
		case `\u`:
			fmt.Println("upd")
		case `\d`:
			fmt.Println("del")
		case `\q`:
			return
		default:
			fmt.Println("Action should be 'list', 'add', 'upd', 'del', or 'quit'")
		}
	}
}
