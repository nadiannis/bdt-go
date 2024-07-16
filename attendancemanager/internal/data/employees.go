package data

type Employee struct {
	ID        string
	Name      string
	IsPresent bool
}

type EmployeeModel struct {
	DB map[string]Employee
}

func (m *EmployeeModel) GetAll() []Employee {
	employees := make([]Employee, 0)
	for _, employee := range m.DB {
		employees = append(employees, employee)
	}
	return employees
}

func (m *EmployeeModel) Add(employee Employee) Employee {
	m.DB[employee.ID] = employee
	return employee
}

func (m *EmployeeModel) UpdatePresenceStatus(employeeID string, isPresent bool) error {
	if employee, exists := m.DB[employeeID]; exists {
		employee.IsPresent = isPresent
		m.DB[employeeID] = employee
		return nil
	}

	return ErrRecordNotFound
}

func (m *EmployeeModel) DeleteByID(employeeID string) error {
	if _, exists := m.DB[employeeID]; !exists {
		return ErrRecordNotFound
	}

	delete(m.DB, employeeID)
	return nil
}
