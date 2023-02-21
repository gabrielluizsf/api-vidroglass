package repository

import (
	"database/sql"
	"fmt"
	"vidroglass/model"

	_ "github.com/mattn/go-sqlite3"
)

func GetCustomers() ([]model.Customer, error) {

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer db.Close()
	rows, err := db.Query("SELECT * FROM customer")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var CustomerList []model.Customer
	var Customer model.Customer

	for rows.Next() {
		err = rows.Scan(
			&Customer.ID,
			&Customer.Nome,
			&Customer.Telefone,
			&Customer.Address.State,
			&Customer.Address.City,
			&Customer.Address.Street,
			&Customer.Address.Number,
			&Customer.Address.ZipCode)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		CustomerList = append(CustomerList, Customer)
	}

	rows.Close()
	return CustomerList, nil
}

func CreateCustomer(customer model.Customer) (model.Customer, error) {

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	query := "INSERT INTO customer " +
		"(name, phone_number, state, city, street, number, zip_number) " +
		"values(?,?,?,?,?,?,?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return model.Customer{}, err
	}

	_, err = stmt.Exec(customer.Nome,
		customer.Telefone,
		customer.Address.State,
		customer.Address.City,
		customer.Address.Street,
		customer.Address.Number,
		customer.Address.ZipCode)

	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}

func GetCustomerByID(customer_id int) (model.Customer, error) {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	row := db.QueryRow("SELECT * FROM customer WHERE id_customer = ?", customer_id)
	var Customer model.Customer

	err = row.Scan(
		&Customer.ID,
		&Customer.Nome,
		&Customer.Telefone,
		&Customer.Address.State,
		&Customer.Address.City,
		&Customer.Address.Street,
		&Customer.Address.Number,
		&Customer.Address.ZipCode)

	if err != nil {
		fmt.Println(err)
		return model.Customer{}, err
	}

	return Customer, nil
}

func UpdateCustomerByID(customer model.Customer) error {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	stmt, err := db.Prepare("UPDATE customer SET name = ?, phone_number = ? WHERE id_customer = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(customer.Nome, customer.Telefone, customer.ID)

	fmt.Println(res)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCustomerByID(customer_id int) error {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM customer WHERE id_customer = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(customer_id)

	fmt.Println(res)

	if err != nil {
		return err
	}

	return nil
}
