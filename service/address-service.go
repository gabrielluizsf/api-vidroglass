package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
	_ "github.com/mattn/go-sqlite3"
)

type addressService struct {
	Address     models.Address
	AddressList []models.Address
}

func NewAddressService() interfaces.AddressService {
	return &addressService{}
}

func (c *addressService) CreateAddress(address models.Address) (int, error) {

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	stmt, err := db.Prepare("insert into address (id_customer, state, city, street, number, zip_number) values (?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(address.Id_customer, address.State, address.City, address.Street, address.Number, address.Cep)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println(id)

	return int(id), nil
}

func (c *addressService) GetAddress() ([]models.Address, error) {

	c.AddressList = nil

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()

	rows, err := db.Query("SELECT * FROM address")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&c.Address.Id_address,
			&c.Address.Id_customer,
			&c.Address.State,
			&c.Address.City,
			&c.Address.Street,
			&c.Address.Number,
			&c.Address.Cep)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		c.AddressList = append(c.AddressList, c.Address)
	}

	rows.Close()
	return c.AddressList, nil
}

func (c *addressService) GetAddressByID(id_address int) (models.Address, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	row := db.QueryRow("SELECT * FROM address WHERE id_address = ?", id_address)

	err = row.Scan(
		&c.Address.Id_address,
		&c.Address.Id_customer,
		&c.Address.State,
		&c.Address.City,
		&c.Address.Street,
		&c.Address.Number,
		&c.Address.Cep)

	if err != nil {
		fmt.Println(err)
		return c.Address, err
	}
	return c.Address, nil
}

func (c *addressService) UpdateAddress(address models.Address) error {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	stmt, err := db.Prepare("UPDATE address SET state = ?, city = ?, street = ?, number = ?, zip_number = ? WHERE id_address = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(address.State, address.City, address.Street, address.Number, address.Cep, address.Id_address)

	fmt.Println(res)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}
