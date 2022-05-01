package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
	_ "github.com/mattn/go-sqlite3"
)

type productTypeService struct {
	Product_type []models.ProductType
}

func NewProductTypeService() interfaces.ProductTypeService {
	return &productTypeService{}
}

func (c *productTypeService) CreateProductType(product_type models.ProductType) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("insert into tipo_produto (nome, descricao) values (?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(product_type.Nome, product_type.Descricao)

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

func (c *productTypeService) GetProductType() ([]models.ProductType, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	rows, err := db.Query("SELECT * FROM tipo_produto")

	if err != nil {
		fmt.Println(err)
		return c.Product_type, err
	}

	c.Product_type = nil

	var id_product_type int
	var Nome string
	var Descricao string

	for rows.Next() {
		err = rows.Scan(&id_product_type, &Nome, &Descricao)
		if err != nil {
			fmt.Println(err)
			return c.Product_type, err
		}
		c.Product_type = append(c.Product_type, models.ProductType{id_product_type, Nome, Descricao})
	}

	rows.Close()
	return c.Product_type, nil
}

func (c *productTypeService) GetProductTypeByID(id_product_type int) (models.ProductType, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	row := db.QueryRow("SELECT * FROM tipo_produto WHERE id_tipo_produto = ?", id_product_type)

	var product_type models.ProductType

	var nome string
	var descricao string

	err = row.Scan(&id_product_type, &nome, &descricao)
	if err != nil {
		fmt.Println(err)
		return product_type, err
	}

	product_type = models.ProductType{id_product_type, nome, descricao}

	return product_type, nil
}

func (c *productTypeService) UpdateProductType(product_type models.ProductType) error {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("UPDATE tipo_produto SET nome = ?, descricao = ? WHERE id_tipo_produto = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(product_type.Nome, product_type.Descricao, product_type.Id_tipo)

	fmt.Println(res)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}
