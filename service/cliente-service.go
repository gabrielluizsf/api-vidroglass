package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
	_ "github.com/mattn/go-sqlite3"
)

type clienteService struct {
	Cliente []models.Cliente
}

func NewClienteService() interfaces.ClienteService {
	return &clienteService{}
}

func (c *clienteService) Save(cliente models.Cliente) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	stmt, err := db.Prepare("INSERT INTO cliente(id_endereco, nome, cpf, telefone) values(?,?,?, ?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(cliente.Id_endereco, cliente.Nome, cliente.Cpf, cliente.Telefone)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println(id)

	db.Close()

	return int(id), nil

}

func (c *clienteService) FindAll() ([]models.Cliente, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	fmt.Println(err)
	rows, err := db.Query("SELECT * FROM cliente")
	fmt.Println(err)

	fmt.Println("Estou prestes a dar erro")
	if err != nil {
		fmt.Println(err)
		return c.Cliente, err
	}
	fmt.Println("Estou prestes a dar erro")

	c.Cliente = nil
	var id_cliente int
	var id_endereco int
	var nome string
	var cpf string
	var telefone string

	for rows.Next() {
		err = rows.Scan(&id_cliente, &id_endereco, &nome, &cpf, &telefone)
		if err != nil {
			fmt.Println(err)
			return c.Cliente, err
		}
		c.Cliente = append(c.Cliente, models.Cliente{id_cliente, id_endereco, nome, cpf, telefone})
	}

	rows.Close()
	db.Close()
	return c.Cliente, nil
}

func (c *clienteService) GetClientById(id_cliente int) (models.Cliente, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	row := db.QueryRow("SELECT * FROM cliente WHERE id_cliente = ?", id_cliente)

	var cliente models.Cliente
	var id_endereco int
	var nome string
	var cpf string
	var telefone string

	err = row.Scan(&id_cliente, &id_endereco, &nome, &cpf, &telefone)
	if err != nil {
		fmt.Println(err)
		return cliente, err
	}

	cliente = models.Cliente{id_cliente, id_endereco, nome, cpf, telefone}

	return cliente, nil
}

func (c *clienteService) UpdateClientById(cliente models.Cliente) error {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	stmt, err := db.Prepare("UPDATE cliente SET id_endereco = ?, nome = ?, cpf = ?, telefone = ? WHERE id_cliente = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(cliente.Id_endereco, cliente.Nome, cliente.Cpf, cliente.Telefone, cliente.Id_cliente)

	fmt.Println(res)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}
