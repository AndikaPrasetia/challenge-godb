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

	// main menu
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
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			customerMenu()
		case 2:
			serviceMenu()
		case 3:
			orderMenu()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// customer menu
func customerMenu() {
	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("================= Customer Menu ==================")
		fmt.Println("1. Create Customer")
		fmt.Println("2. View Of List Customers")
		fmt.Println("3. View Details Customer by ID")
		fmt.Println("4. Update Customer")
		fmt.Println("5. Delete Customer")
		fmt.Println("6. Back to Main Menu")
		fmt.Println(strings.Repeat("=", 50))

		fmt.Print("Enter your choice: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			createCustomer()
		case 2:
			customers := viewOfListCustomers()
			for _, customer := range customers {
				fmt.Println(customer.Id, customer.Name, customer.Phone, customer.Address, customer.CreatedAt, customer.UpdatedAt)
			}
		case 3:
			viewDetailCustomerById()
		case 4:
			updateCustomer()
		case 5:
			deleteCustomer()
		case 6:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// service menu
func serviceMenu() {
	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("================== Service Menu ==================")
		fmt.Println("1. Create Service")
		fmt.Println("2. View Of List Services")
		fmt.Println("3. View Details Service by ID")
		fmt.Println("4. Update Service")
		fmt.Println("5. Delete Service")
		fmt.Println("6. Back to Main Menu")
		fmt.Println(strings.Repeat("=", 50))

		fmt.Print("Enter your choice: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			createService()
		case 2:
			services := viewOfListServices()
			for _, service := range services {
				fmt.Println(service.Id, service.Name, service.Unit, service.Price, service.CreatedAt, service.UpdatedAt)
			}
		case 3:
			viewDetailServiceById()
		case 4:
			updateCustomer()
		case 5:
			deleteCustomer()
		case 6:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// order menu
func orderMenu() {
	for {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("=================== Order Menu ===================")
		fmt.Println("1. Create Order")
		fmt.Println("2. View Of List Orders")
		fmt.Println("3. View Details Order by ID")
		fmt.Println("4. Update Order")
		fmt.Println("5. Delete Order")
		fmt.Println("6. Back to Main Menu")
		fmt.Println(strings.Repeat("=", 50))

		fmt.Print("Enter your choice: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			createCustomer()
		case 2:
			customers := viewOfListCustomers()
			for _, customer := range customers {
				fmt.Println(customer.Id, customer.Name, customer.Phone, customer.Address, customer.CreatedAt, customer.UpdatedAt)
			}
		case 3:
			viewDetailCustomerById()
		case 4:
			updateCustomer()
		case 5:
			deleteCustomer()
		case 6:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// create customer
func createCustomer() {
	db := connectDb()
	defer db.Close()
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	customer := entity.CustomerEnrollment{}

	fmt.Println("Enter Customer Details:")

	fmt.Print("Customer ID: ")
	scanner.Scan()
	customer.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Name: ")
	scanner.Scan()
	customer.Name = scanner.Text()

	fmt.Print("Phone: ")
	scanner.Scan()
	customer.Phone = scanner.Text()

	fmt.Print("Address: ")
	scanner.Scan()
	customer.Address = scanner.Text()

	sqlStatement := "INSERT INTO customer (customer_id, name, phone, address) VALUES ($1, $2, $3, $4);"

	_, err = db.Exec(sqlStatement, customer.Id, customer.Name, customer.Phone, customer.Address)

	if err != nil {
		fmt.Println("Customer ID already exists. Please enter a different ID")
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

// view customers
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

// view customer by id
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

// update customer
func updateCustomer() {
	db := connectDb()
	defer db.Close()
	var err error

	customer := entity.CustomerEnrollment{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Customer ID: ")
	scanner.Scan()
	customer.Id, _ = strconv.Atoi(scanner.Text())

	sqlCheck := "SELECT customer_id FROM customer WHERE customer_id = $1"
	err = db.QueryRow(sqlCheck, customer.Id).Scan(&customer.Id)

	if err == sql.ErrNoRows {
		fmt.Println("Customer not found.")
	} else if err != nil {
		fmt.Println("Error checking customer ID:", err)
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

	sqlStatement := "UPDATE customer SET name = $2, phone = $3, address = $4 WHERE customer_id = $1;"

	_, err = db.Exec(sqlStatement, customer.Id, customer.Name, customer.Phone, customer.Address)

	if err != nil {
		fmt.Println("Error Udate Data", err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

// delete customer
func deleteCustomer() {
	db := connectDb()
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	var customer_id int

	fmt.Print("Enter Customer ID: ")
	scanner.Scan()
	customer_id, _ = strconv.Atoi(scanner.Text())

	sqlCheckCustomer := "SELECT customer_id FROM customer WHERE customer_id = $1;"

	err := db.QueryRow(sqlCheckCustomer, customer_id).Scan(&customer_id)

	if err == sql.ErrNoRows {
		fmt.Println("Customer ID not found. Please enter a different ID")

	} else if err != nil {
		fmt.Println("Error checking customer ID:", err)
	}

	sqlCheckOrder := "SELECT customer_id FROM customer WHERE customer_id = $1;"
	var order_id int

	err = db.QueryRow(sqlCheckOrder, order_id).Scan(&order_id)

	if err == nil {
		fmt.Println("Customer ID is being used in orders. Please delete the order first.")

	} else if err != sql.ErrNoRows {
		fmt.Println("Error checking orders ID:", err)
	}

	sqlStatement := "DELETE FROM customer WHERE customer_id = $1;"
	_, err = db.Exec(sqlStatement, customer_id)

	if err != nil {
		fmt.Println("Error deleting customer:", err)
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}

// create service
func createService() {
	db := connectDb()
	defer db.Close()
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	service := entity.ServiceEnrollment{}

	fmt.Println("Enter Service Details:")

	fmt.Print("Service ID: ")
	scanner.Scan()
	service.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Name: ")
	scanner.Scan()
	service.Name = scanner.Text()

	fmt.Print("Unit: ")
	scanner.Scan()
	service.Unit = scanner.Text()

	fmt.Print("Price: ")
	scanner.Scan()
	service.Price, _ = strconv.Atoi(scanner.Text())

	sqlStatement := "INSERT INTO service (service_id, service_name, unit, price) VALUES ($1, $2, $3, $4);"

	_, err = db.Exec(sqlStatement, service.Id, service.Name, service.Unit, service.Price)

	if err != nil {
		fmt.Println("Customer ID already exists. Please enter a different ID")
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

// view services
func viewOfListServices() []entity.ServiceEnrollment {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM service;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	services := scanService(rows)
	return services
}

// view service by id
func viewDetailServiceById() {
	db := connectDb()
	defer db.Close()
	var err error

	service := entity.ServiceEnrollment{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Customer Id: ")
	scanner.Scan()
	service.Id, _ = strconv.Atoi(scanner.Text())

	sqlStatement := "SELECT * FROM service WHERE service_id = $1;"

	err = db.QueryRow(sqlStatement, service.Id).Scan(&service.Id, &service.Name, &service.Unit, &service.Price, &service.CreatedAt, &service.UpdatedAt)
	if err != nil {
		fmt.Println("Customer not found.")
	} else {
		fmt.Println(service)
	}
}

// =============== HELPER FUNCTION ===============

// scan customer
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

// scan service
func scanService(rows *sql.Rows) []entity.ServiceEnrollment {
	services := []entity.ServiceEnrollment{}
	var err error

	for rows.Next() {
		service := entity.ServiceEnrollment{}
		err := rows.Scan(&service.Id, &service.Name, &service.Unit, &service.Price, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			panic(err)
		}
		services = append(services, service)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return services
}
