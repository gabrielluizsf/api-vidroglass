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
	ProductType     models.ProductType
	ProductTypeList []models.ProductType
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
		return c.ProductTypeList, err
	}

	c.ProductTypeList = nil

	for rows.Next() {
		err = rows.Scan(
			&c.ProductType.Id_tipo,
			&c.ProductType.Nome,
			&c.ProductType.Descricao)

		if err != nil {
			fmt.Println(err)
			return c.ProductTypeList, err
		}
		c.ProductTypeList = append(c.ProductTypeList, c.ProductType)
	}

	rows.Close()
	return c.ProductTypeList, nil
}

func (c *productTypeService) GetProductTypeByID(id_product_type int) (models.ProductType, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	row := db.QueryRow("SELECT * FROM tipo_produto WHERE id_tipo_produto = ?", id_product_type)

	var product_type models.ProductType

	err = row.Scan(
		&c.ProductType.Id_tipo,
		&c.ProductType.Nome,
		&c.ProductType.Descricao)

	if err != nil {
		fmt.Println(err)
		return product_type, err
	}

	return c.ProductType, nil
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
