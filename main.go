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

	sqlStatement := "INSERT INTO customer (customer_id, name, phone, address) VALUES ($1, $2, $3, $4);"

	_, err = db.Exec(sqlStatement, customerEnrollment.Id, customerEnrollment.Name, customerEnrollment.Phone, customerEnrollment.Address)

	if err != nil {
		fmt.Println("Customer ID already exists. Please enter a different ID")
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

// view customer
func viewOfListCustomers() []entity.CustomerEnrollment {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM customer;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := scanCustomer(rows)
	return customers
}

func scanCustomer(rows *sql.Rows) []entity.CustomerEnrollment {
	customers := []entity.CustomerEnrollment{}
	var err error

	for rows.Next() {
		customer := entity.CustomerEnrollment{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			panic(err)
		}
		customers = append(customers, customer)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return customers
}

func viewDetailCustomerById() {
	db := connectDb()
	defer db.Close()
	var err error

	customer := entity.CustomerEnrollment{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Customer Id: ")
	scanner.Scan()
	customer.Id, _ = strconv.Atoi(scanner.Text())

	sqlStatement := "SELECT * FROM customer WHERE customer_id = $1;"

	err = db.QueryRow(sqlStatement, customer.Id).Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		fmt.Println("Customer not found.")
	} else {
		fmt.Println(customer)
	}
}

func updateCustomer() {
	db := connectDb()
	defer db.Close()
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	customer := entity.CustomerEnrollment{}

	fmt.Println("Enter Customer Details:")

	fmt.Print("Customer ID: ")
	scanner.Scan()
	customer.Id, _ = strconv.Atoi(scanner.Text())

	sqlCheck := "SELECT customer_id FROM customer WHERE customer_id = $1"
	err = db.QueryRow(sqlCheck, customer.Id).Scan(&customer.Id)

	if err == sql.ErrNoRows {
		fmt.Println("Customer not found.")
		return
	} else if err != nil {
		fmt.Println("Error checking customer ID:", err)
		return
	}

	fmt.Print("Name: ")
	scanner.Scan()
	customer.Name = scanner.Text()

	fmt.Print("Phone: ")
	scanner.Scan()
	customer.Phone = scanner.Text()

	fmt.Print("Address: ")
	scanner.Scan()
	customer.Address = scanner.Text()

	sqlStatement := "UPDATE customer SET name = $2, phone = $3, address = $4, created_at = $5, updated_at = $6 WHERE customer_id = $1;"

	_, err = db.Exec(sqlStatement, customer.Id, customer.Name, customer.Phone, customer.Address, customer.CreatedAt, customer.UpdatedAt)

	if err != nil {
		fmt.Println("Error Udate Data", err)
	} else {
		fmt.Println("Successfully Update Data!")
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
		case "2":
			customers := viewOfListCustomers()
			for _, customer := range customers {
				fmt.Println(customer.Id, customer.Name, customer.Phone, customer.Address, customer.CreatedAt, customer.UpdatedAt)
			}
		case "3":
			viewDetailCustomerById()
		case "4":
			updateCustomer()
		// case "5":
		// 	deleteCustomer()
		case "6":
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
