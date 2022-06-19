package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
	_ "github.com/mattn/go-sqlite3"
)

type productService struct {
	Product []models.Product
}

func NewProductService() interfaces.ProductService {
	return &productService{}
}

func (c *productService) CreateProduct(product models.Product) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("insert into produto (id_tipo, valor_metragem, espessura, cor) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(product.Id_tipo, product.Valor_metragem, product.Espessura, product.Cor)

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

func (c *productService) GetProduct() ([]models.Product, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	rows, err := db.Query("SELECT * FROM produto")

	if err != nil {
		fmt.Println(err)
		return c.Product, err
	}

	c.Product = nil

	var id_produto int
	var id_tipo int
	var valor_metragem float32
	var espessura float32
	var cor string

	for rows.Next() {
		err = rows.Scan(&id_produto, &id_tipo, &valor_metragem, &espessura, &cor)
		if err != nil {
			fmt.Println(err)
			return c.Product, err
		}
		c.Product = append(c.Product, models.Product{
			Id_produto:     id_produto,
			Id_tipo:        id_tipo,
			Valor_metragem: valor_metragem,
			Espessura:      espessura,
			Cor:            cor})
	}

	rows.Close()
	return c.Product, nil
}

func (c *productService) GetProductByID(id_produto int) (models.Product, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	row := db.QueryRow("SELECT * FROM produto WHERE id_produto = ?", id_produto)

	var product models.Product
	var id_tipo int
	var valor_metragem float32
	var espessura float32
	var cor string

	err = row.Scan(&id_produto, &id_tipo, &valor_metragem, &espessura, &cor)
	if err != nil {
		fmt.Println(err)
		return product, err
	}

	product = models.Product{
		Id_produto:     id_produto,
		Id_tipo:        id_tipo,
		Valor_metragem: valor_metragem,
		Espessura:      espessura,
		Cor:            cor}

	return product, nil
}

func (c *productService) UpdateProduct(product models.Product) error {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	stmt, err := db.Prepare("UPDATE produto SET id_tipo = ?, valor_metragem = ?, espessura = ?, cor = ? WHERE id_produto = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(product.Id_tipo, product.Valor_metragem, product.Espessura, product.Cor, product.Id_produto)

	fmt.Println(res)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}
