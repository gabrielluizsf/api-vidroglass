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
	Address []models.Address
}

func NewAddressService() interfaces.AddressService {
	return &addressService{}
}

func (c *addressService) CreateAddress(address models.Address) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("insert into endereco (rua, numero, cep, cidade, estado) values (?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(address.Street, address.Number, address.Cep, address.State)

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

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	rows, err := db.Query("SELECT * FROM endereco")

	if err != nil {
		fmt.Println(err)
		return c.Address, err
	}

	c.Address = nil

	var id_endereco int
	var rua string
	var numero int
	var cep string
	var cidade string
	var estado string

	for rows.Next() {
		err = rows.Scan(&id_endereco, &rua, &numero, &cep, &cidade, &estado)
		if err != nil {
			fmt.Println(err)
			return c.Address, err
		}
		c.Address = append(c.Address, models.Address{id_endereco, rua, numero, cep, cidade, estado})
	}

	rows.Close()
	return c.Address, nil
}

func (c *addressService) GetAddressByID(id_cliente int) (models.Address, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	row := db.QueryRow("SELECT * FROM endereco WHERE id_endereco = ?", id_cliente)

	var address models.Address
	var id_endereco int
	var rua string
	var numero int
	var cep string
	var cidade string
	var estado string

	err = row.Scan(&id_endereco, &rua, &numero, &cep, &cidade, &estado)
	if err != nil {
		fmt.Println(err)
		return address, err
	}

	address = models.Address{id_endereco, rua, numero, cep, cidade, estado}

	return address, nil
}

func (c *addressService) UpdateAddress(address models.Address) error {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	stmt, err := db.Prepare("UPDATE endereco SET rua = ?, numero = ?, cep = ?, cidade = ?, estado = ? WHERE id_endereco = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(address.Street, address.Number, address.Cep, address.City, address.State, address.Id_address)

	fmt.Println(res)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}
