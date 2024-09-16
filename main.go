package main

import (
	"bufio"
	"challenge-godb/entity"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "6521"
	dbname   = "enigma_laundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Connected!")
	}

	return db
}

func main() {
	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("=================== Main Menu ====================")
		fmt.Println("1. Customer")
		fmt.Println("2. Service")
		fmt.Println("3. Order")
		fmt.Println("4. Exit")
		fmt.Println(strings.Repeat("=", 50))

		fmt.Print("Enter your choice: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			customerMenu()
		// case "2":
		// 	serviceMenu()
		// case "3":
		// 	orderMenu()
		case "4":
			fmt.Println("Exiting application...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// membuat customer
func createCustomer() {
	db := connectDb()
	defer db.Close()
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	customerEnrollment := entity.CustomerEnrollment{}

	fmt.Println("Enter Customer Details:")

	fmt.Print("Customer ID: ")
	scanner.Scan()
	customerEnrollment.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Name: ")
	scanner.Scan()
	customerEnrollment.Name = scanner.Text()

	fmt.Print("Phone: ")
	scanner.Scan()
	customerEnrollment.Phone = scanner.Text()

	fmt.Print("Address: ")
	scanner.Scan()
	customerEnrollment.Address = scanner.Text()

	sqlStatement := "INSERT INTO customer (customer_id, name, phone, address) VALUES ($1, $2, $3, $4)"

	_, err = db.Exec(sqlStatement, customerEnrollment.Id, customerEnrollment.Name, customerEnrollment.Phone, customerEnrollment.Address)

	if err != nil {
		fmt.Println("Customer ID already exists. Please enter a different ID")
	} else {
		fmt.Println("Successfully Inser Data!")
	}

}

// menu customer
func customerMenu() {
	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("================= Customer Menu ==================")
		fmt.Println("1. Create Customer")
		fmt.Println("2. View List of Customers")
		fmt.Println("3. View Details Customer by ID")
		fmt.Println("4. Update Customer")
		fmt.Println("5. Delete Customer")
		fmt.Println("6. Back to Main Menu")
		fmt.Println(strings.Repeat("=", 50))

		fmt.Print("Enter your choice: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			createCustomer()
		// case "2":
		// 	viewCustomers()
		// case "3":
		// 	viewCustomerById()
		// case "4":
		// 	updateCustomer()
		// case "5":
		// 	deleteCustomer()
		case "6":
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
