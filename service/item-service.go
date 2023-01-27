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
	Item     models.Item
	ItemList []models.Item
}

func NewItemService() interfaces.ItemService {
	return &itemService{}
}

func (c *itemService) Save(item models.Item) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))

	total_value, err := c.getProductValue(item.Id_produto)
	if err != nil {
		return 0, err
	}
	stmt, err := db.Prepare("INSERT INTO item(id_invoice, id_product, amount, quantity, discount, metreage) values(?,?,?, ?,?,?)")
	if err != nil {
		return 0, err
	}

	total_value *= item.Metragem
	res, err := stmt.Exec(item.Id_nota, item.Id_produto, total_value, item.Quantidade, item.Desconto, item.Metragem)

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
	query := fmt.Sprintf("SELECT " +
		"i.id_item, " +
		"i.id_invoice, " +
		"p.color, " +
		"tp.name, " +
		"i.discount, " +
		"i.metreage, " +
		"i.quantity, " +
		"i.amount " +
		"FROM " +
		"item i " +
		"JOIN " +
		"product as p on i.id_product = p.id_product " +
		"JOIN " +
		"product_type as tp on p.id_type = tp.id_type ")
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
		return c.ItemList, err
	}

	c.ItemList = nil

	for rows.Next() {
		err = rows.Scan(
			&c.Item.Id_item,
			&c.Item.Id_nota,
			&c.Item.Cor,
			&c.Item.Nome,
			&c.Item.Desconto,
			&c.Item.Metragem,
			&c.Item.Quantidade,
			&c.Item.Valor)

		if err != nil {
			fmt.Println(err)
			return c.ItemList, err
		}
		c.ItemList = append(c.ItemList, c.Item)
	}

	rows.Close()
	db.Close()
	return c.ItemList, nil
}

func (c *itemService) GetItemById(id_item int) (models.Item, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))

	query := fmt.Sprintf("SELECT " +
		"i.id_item, " +
		"i.id_invoice, " +
		"p.color, " +
		"tp.name, " +
		"i.discount, " +
		"i.metreage, " +
		"i.quantity, " +
		"i.amount " +
		"FROM " +
		"item i " +
		"JOIN " +
		"product as p on i.id_product = p.id_product " +
		"JOIN " +
		"product_type as tp on p.id_type = tp.id_type " +
		"WHERE " +
		"i.id_item = ?")
	row := db.QueryRow(query, id_item)

	err = row.Scan(
		&c.Item.Id_item,
		&c.Item.Id_nota,
		&c.Item.Cor,
		&c.Item.Nome,
		&c.Item.Desconto,
		&c.Item.Metragem,
		&c.Item.Quantidade,
		&c.Item.Valor)

	if err != nil {
		fmt.Println(err)
		return c.Item, err
	}

	return c.Item, nil
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

func (c *itemService) getProductValue(id_produto int) (float64, error) {

	var totalValue float64

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	row := db.QueryRow("SELECT value_per_meter FROM product where id_product = ?", id_produto)

	if err != nil {
		return 0, err
	}

	err = row.Scan(&totalValue)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return totalValue, nil

}
func (c *itemService) GetInvoiceItens(id_nota int) ([]models.ItemPayload, error) {

	var itemPayload models.ItemPayload
	var itemPayloads []models.ItemPayload
	fmt.Print(id_nota)

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	rows, err := db.Query("SELECT "+
		"i.quantity, i.amount, i.discount,"+
		"i.metreage, p.thickness, p.color, tp.name "+
		"FROM item as i "+
		"join product as p on p.id_product = i.id_product "+
		"join product_type as tp on tp.id_type = p.id_type "+
		"where i.id_invoice = ?", id_nota)

	if err != nil {
		fmt.Println(err)
		return itemPayloads, err
	}

	for rows.Next() {
		err = rows.Scan(&itemPayload.Quantidade, &itemPayload.Valor,
			&itemPayload.Desconto, &itemPayload.Metragem,
			&itemPayload.Espessura, &itemPayload.Cor, &itemPayload.Nome)
		if err != nil {
			fmt.Println(err)
			return itemPayloads, err
		}
		itemPayloads = append(itemPayloads, itemPayload)
	}

	rows.Close()
	db.Close()

	return itemPayloads, nil
}
