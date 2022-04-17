package service

import (
	"github.com/mariarobertap/api-vidroglass/models"
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)


type CustomerService interface {

	Save(models.Customer) 
	FindAll() ([]models.Customer, error)

}

type customerService struct {
	customer []models.Customer
}


func NewCustomerService() CustomerService{
	return &customerService{}
}

func (c *customerService) Save(customer models.Customer){
	

	db, err := sql.Open("sqlite3", "./database/data.db")
	stmt, err := db.Prepare("INSERT INTO customers(first_name, last_name) values(?,?)")
	checkErr(err)

	res, err := stmt.Exec(customer.FirstName, customer.LastName)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	db.Close()

}

func (c *customerService) FindAll() ([]models.Customer, error){

	db, err := sql.Open("sqlite3", "./database/data.db")
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
        fmt.Println(err)
		return c.customer, err
    }
	c.customer = nil
	var uid int
	var fisrtName string
	var lastName string


	for rows.Next() {
		err = rows.Scan(&uid, &fisrtName, &lastName)
		if err != nil {
			fmt.Println(err)
			return c.customer, err
		}
		c.customer = append(c.customer, models.Customer{fisrtName, lastName})
	}

	rows.Close() 
	db.Close()
	return c.customer, nil
}


func checkErr(err error){
	if err != nil {
        fmt.Println(err)
    }
}