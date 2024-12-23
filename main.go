package main

import (
	"fmt"

	"github.com/NurNurzhanuly/Assignment1/Assignment1/bank"
	"github.com/NurNurzhanuly/Assignment1/Assignment1/employees"
	"github.com/NurNurzhanuly/Assignment1/Assignment1/library"
	"github.com/NurNurzhanuly/Assignment1/Assignment1/shapes"
)

func main() {
	fmt.Println("Welcome to the Assignment 1 Application!")
	fmt.Println("Please select a module to work with:")

	for {
		fmt.Println("1. Library Management")
		fmt.Println("2. Shapes Calculator")
		fmt.Println("3. Employee Management")
		fmt.Println("4. Bank Account Management")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			manageLibrary()
		case 2:
			manageShapes()
		case 3:
			manageEmployees()
		case 4:
			manageBankAccounts()
		case 5:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func manageLibrary() {
	lib := library.NewLibrary()
	for {
		fmt.Println("Library Management:")
		fmt.Println("1. Add a Book")
		fmt.Println("2. Borrow a Book")
		fmt.Println("3. Return a Book")
		fmt.Println("4. List Books")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, title, author string
			fmt.Print("Enter Book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Book Title: ")
			fmt.Scan(&title)
			fmt.Print("Enter Book Author: ")
			fmt.Scan(&author)
			lib.AddBook(library.Book{ID: id, Title: title, Author: author, IsBorrowed: false})
		case 2:
			var id string
			fmt.Print("Enter Book ID to Borrow: ")
			fmt.Scan(&id)
			lib.BorrowBook(id)
		case 3:
			var id string
			fmt.Print("Enter Book ID to Return: ")
			fmt.Scan(&id)
			lib.ReturnBook(id)
		case 4:
			lib.ListBooks()
		case 5:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func manageShapes() {
	shapesList := []shapes.Shape{
		shapes.NewRectangle(5, 3),
		shapes.NewCircle(4),
		shapes.NewSquare(6),
		shapes.NewTriangle(3, 4, 5),
	}

	fmt.Println("Shapes Calculator:")
	for _, shape := range shapesList {
		shapes.PrintShapeDetails(shape)
	}
}

func manageEmployees() {
	company := employees.NewCompany()
	for {
		fmt.Println("Employee Management:")
		fmt.Println("1. Add Full-Time Employee")
		fmt.Println("2. Add Part-Time Employee")
		fmt.Println("3. List Employees")
		fmt.Println("4. Back to Main Menu")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id uint64
			var name string
			var salary uint32
			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Employee Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Salary: ")
			fmt.Scan(&salary)
			company.AddEmployee(employees.NewFullTimeEmployee(id, name, salary))
		case 2:
			var id uint64
			var name string
			var hourlyRate uint64
			var hoursWorked float32
			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Employee Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Hourly Rate: ")
			fmt.Scan(&hourlyRate)
			fmt.Print("Enter Hours Worked: ")
			fmt.Scan(&hoursWorked)
			company.AddEmployee(employees.NewPartTimeEmployee(id, name, hourlyRate, hoursWorked))
		case 3:
			company.ListEmployees()
		case 4:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func manageBankAccounts() {
	account := bank.NewBankAccount("12345", "John Doe", 1000.0)
	for {
		fmt.Println("Bank Account Management:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get Balance")
		fmt.Println("4. Back to Main Menu")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var amount float64
			fmt.Print("Enter Deposit Amount: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			var amount float64
			fmt.Print("Enter Withdrawal Amount: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			fmt.Printf("Current Balance: %.2f\n", account.GetBalance())
		case 4:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
