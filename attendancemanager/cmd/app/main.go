package main

import (
	"attendancemanager/internal/data"
	"fmt"

	"github.com/google/uuid"
)

type application struct {
	models data.Models
}

func main() {
	app := application{
		models: *data.NewModels(),
	}

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
	fmt.Println(app.models.Employees.GetAll())
	app.models.Employees.Add(employee2)
	fmt.Println(app.models.Employees.GetAll())

	err := app.models.Employees.UpdatePresenceStatus(employee1.ID, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(app.models.Employees.GetAll())

	err = app.models.Employees.DeleteByID(employee2.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(app.models.Employees.GetAll())
}
