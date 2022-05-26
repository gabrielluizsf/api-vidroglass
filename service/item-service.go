package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
	_ "github.com/mattn/go-sqlite3"
)

type itemService struct {
	Item []models.Item
}

func NewItemService() interfaces.ItemService {
	return &itemService{}
}

func (c *itemService) Save(item models.Item) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	stmt, err := db.Prepare("INSERT INTO item(id_produto, id_nota, quantidade, valor, desconto, metragem_produto) values(?,?,?, ?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(item.Id_produto, item.Id_nota, item.Quantidade, item.Valor, item.Desconto, item.Metragem)

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

func (c *itemService) FindAll() ([]models.Item, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	rows, err := db.Query("SELECT * FROM item")

	if err != nil {
		fmt.Println(err)
		return c.Item, err
	}

	c.Item = nil
	var Id_item int
	var Id_produto int
	var Id_nota string
	var Valor float64
	var Quantidade int
	var Desconto int
	var Metragem float64

	for rows.Next() {
		err = rows.Scan(&Id_item, &Id_produto, &Id_nota, &Valor, &Quantidade, &Desconto, &Metragem)
		if err != nil {
			fmt.Println(err)
			return c.Item, err
		}
		c.Item = append(c.Item, models.Item{Id_item, Id_produto, Id_nota, Valor, Quantidade, Desconto, Metragem})
	}

	rows.Close()
	db.Close()
	return c.Item, nil
}

func (c *itemService) GetItemById(id_item int) (models.Item, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	row := db.QueryRow("SELECT * FROM item WHERE id_item = ?", id_item)

	var item models.Item
	var Id_item int
	var Id_produto int
	var Id_nota string
	var Valor float64
	var Quantidade int
	var Desconto int
	var Metragem float64

	err = row.Scan(&Id_item, &Id_produto, &Id_nota, &Valor, &Quantidade, &Desconto, &Metragem)
	if err != nil {
		fmt.Println(err)
		return item, err
	}

	item = models.Item{Id_item, Id_produto, Id_nota, Valor, Quantidade, Desconto, Metragem}

	return item, nil
}

func (c *itemService) UpdateItemById(item models.Item) error {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	stmt, err := db.Prepare("UPDATE item SET id_produto = ?, id_nota = ?, quantidade = ?, valor = ? desconto = ? metragem_produto = ? WHERE id_item = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(item.Id_produto, item.Id_nota, item.Quantidade, item.Valor, item.Desconto, item.Metragem)

	fmt.Println(res)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}
