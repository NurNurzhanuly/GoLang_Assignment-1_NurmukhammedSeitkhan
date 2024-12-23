package employees

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("FullTime - ID: %d, Name: %s, Salary: %d", f.ID, f.Name, f.Salary)
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("PartTime - ID: %d, Name: %s, Earnings: %.2f", p.ID, p.Name, float64(p.HourlyRate)*float64(p.HoursWorked))
}

type Company struct {
	Employees map[string]Employee
}

func (c *Company) AddEmployee(id string, emp Employee) {
	c.Employees[id] = emp
}

func (c *Company) ListEmployees() {
	for _, emp := range c.Employees {
		fmt.Println(emp.GetDetails())
	}
}
