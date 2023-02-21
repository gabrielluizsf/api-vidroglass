package repository

import (
	"database/sql"
	"fmt"
	"vidroglass/model"
)

func CreateProduct(product model.Product) (int, error) {

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	stmt, err := db.Prepare("insert into product (id_type, value_per_meter, total_value, thickness, color) values (?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(product.ProductTypeID, product.ValuePerMeter, product.TotalValue, product.Espessura, product.Color)

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

func GetProducts() ([]model.ProductObject, error) {

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	rows, err := db.Query(`SELECT 
								p.id_product,
								t.name,
								p.value_per_meter,
								p.total_value,
								p.thickness,
								p.color
							FROM 
								product p
							JOIN 
								product_type t on p.id_type = t.id_type
								`)

	if err != nil {
		fmt.Println(err)
		return []model.ProductObject{}, err
	}

	var ProductObjectList []model.ProductObject
	var ProductObject model.ProductObject
	for rows.Next() {
		err = rows.Scan(
			&ProductObject.ProductID,
			&ProductObject.ProductType,
			&ProductObject.ValuePerMeter,
			&ProductObject.TotalValue,
			&ProductObject.Espessura,
			&ProductObject.Color)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		ProductObjectList = append(ProductObjectList, ProductObject)
	}

	rows.Close()

	return ProductObjectList, nil
}

func GetProductByID(id_produto int) (model.ProductObject, error) {
	db, err := sql.Open("sqlite3", "database.db")
	row := db.QueryRow("SELECT "+
		"p.id_product, "+
		"t.name, "+
		"p.value_per_meter, "+
		"p.total_value, "+
		"p.thickness, "+
		"p.color "+
		"FROM "+
		"product p "+
		"JOIN "+
		"product_type t on p.id_type = t.id_type "+
		"WHERE p.id_product = ?", id_produto)

	var product model.ProductObject

	err = row.Scan(
		&product.ProductID,
		&product.ProductType,
		&product.ValuePerMeter,
		&product.TotalValue,
		&product.Espessura,
		&product.Color)

	if err != nil {
		//LOG
		fmt.Println(err)
		return model.ProductObject{}, err
	}

	return product, nil
}

func UpdateProduct(product model.Product) error {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	stmt, err := db.Prepare("UPDATE product SET id_tipo = ?, valor_metragem = ?, espessura = ?, cor = ? WHERE id_produto = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.ProductTypeID, product.ValuePerMeter, product.TotalValue, product.Espessura, product.Color, product.ProductID)

	if err != nil {
		return err
	}
	return nil
}

func DeleteProductById(id_product int) error {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	query, err := db.Prepare("DELETE FROM product WHERE id_product = ?")

	if err != nil {
		return err
	}

	res, err := query.Exec(id_product)

	fmt.Println(res)
	if err != nil {
		return err
	}

	return nil
}

func getProductValue(product_id int) (model.Product, error) {
	db, err := sql.Open("sqlite3", "database.db")
	row := db.QueryRow("SELECT "+
		"value_per_meter, "+
		"total_value "+
		"FROM "+
		"product "+
		"WHERE id_product = ?", product_id)

	var product model.Product

	err = row.Scan(
		&product.ValuePerMeter,
		&product.TotalValue)

	if err != nil {
		//LOG
		fmt.Println(err)
		return model.Product{}, err
	}

	return product, nil
}

func CreateProductType(product_type model.ProductType) (int, error) {

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	stmt, err := db.Prepare("insert into product_type (name, description) values (?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(product_type.Name, product_type.Description)

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

func GetProductTypes() ([]model.ProductType, error) {

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	rows, err := db.Query("SELECT * FROM product_type")

	var productTypeList []model.ProductType
	var productType model.ProductType

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&productType.TypeID,
			&productType.Name,
			&productType.Description)

		if err != nil {
			fmt.Println(err)
			return productTypeList, err
		}
		productTypeList = append(productTypeList, productType)
	}

	rows.Close()
	return productTypeList, nil
}

func GetProductTypeByID(id_product_type int) (model.ProductType, error) {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	row := db.QueryRow("SELECT * FROM product_type WHERE id_type = ?", id_product_type)

	var product_type model.ProductType

	err = row.Scan(
		&product_type.TypeID,
		&product_type.Name,
		&product_type.Description)

	if err != nil {
		fmt.Println(err)
		return product_type, err
	}

	return product_type, nil
}

func UpdateProductType(product_type model.ProductType) error {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	stmt, err := db.Prepare("UPDATE product_type SET name = ?, description = ? WHERE id_type = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(product_type.Name, product_type.Description, product_type.TypeID)

	fmt.Println(res)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProductTypeByID(id_tipo int) error {

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	query, err := db.Prepare("DELETE FROM product_type WHERE id_type = ?")

	if err != nil {
		return err
	}

	_, err = query.Exec(id_tipo)

	if err != nil {
		return err
	}

	return nil

}
