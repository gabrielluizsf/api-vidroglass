package service

import (
	"github.com/mariarobertap/api-vidroglass/models"
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)


type ClienteService interface {


	FindAll() ([]models.Cliente, error)
	Save(models.Cliente) (int, error)


}

type clienteService struct {
	Cliente []models.Cliente
}


func NewClienteService() ClienteService{
	return &clienteService{}
}

func (c *clienteService) Save(cliente models.Cliente) (int, error){
	

	db, err := sql.Open("sqlite3", "./database/data.db")
	stmt, err := db.Prepare("INSERT INTO cliente(id_endereco, nome, cpf, telefone) values(?,?,?, ?)")
	if(err != nil){
		return 0, err
	}

	res, err := stmt.Exec(cliente.Id_endereco, cliente.Nome, cliente.Cpf, cliente.Telefone)
	
	if(err != nil){
		return 0, err
	}

	id, err := res.LastInsertId()
	if(err != nil){
		return 0, err
	}

	fmt.Println(id)

	db.Close()

	return int(id), nil

}


func (c *clienteService) FindAll() ([]models.Cliente, error){

	db, err := sql.Open("sqlite3", "./database/data.db")
	rows, err := db.Query("SELECT * FROM cliente")
	if err != nil {
        fmt.Println(err)
		return c.Cliente, err
    }
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

