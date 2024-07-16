package main

import (
	"attendancemanager/internal/data"
	"attendancemanager/internal/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
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

	fmt.Println(`=============================== ATTENDANCE MANAGER ===============================`)
	displayCommands()

	for {
		input := utils.GetInput(app.scanner, "\n> ")
		if input == "" {
			fmt.Println("input is required")
			continue
		}

		parts := strings.Fields(input)
		action := parts[0]

		switch action {
		case `\l`:
			app.getAllEmployees(parts)
		case `\a`:
			app.addEmployee(parts)
		case `\u`:
			app.updateEmployeeStatus(parts)
		case `\d`:
			app.deleteEmployee(parts)
		case `\c`:
			if len(parts) != 1 {
				fmt.Println(`input should be \a`)
				continue
			}
			displayCommands()
		case `\q`:
			fmt.Println("bye!")
			return
		default:
			fmt.Println(`action should be '\l', '\a', '\u', '\d', '\c', or '\q'`)
		}
	}
}

func displayCommands() {
	fmt.Println(`
	Commands:
	\l                              => Get list of employees
	\a                              => Add a new employee
	\u [employee-id] [status (y/n)] => Update employee presence status
	\d [employee-id]                => Delete an employee
	\c                              => Show all commands
	\q                              => Quit`)
}
