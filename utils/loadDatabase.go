package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"vidroglass/models"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type Invoice struct {
	InvoiceList []models.Invoice `json:"invoice"`
}
type Item struct {
	ItemList []models.Item `json:"item"`
}

type Customers struct {
	CustomerList []models.Cliente `json:"customer"`
}

type Address struct {
	AddressList []models.Address `json:"address"`
}

type ProductType struct {
	ProductTypeList []models.ProductType `json:"product_type"`
}

type Product struct {
	ProductList []models.Product `json:"product"`
}

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if os.Getenv("LOAD_CUSTOMER") == "S" {
		err = loadCustomersTable()
		if err != nil {
			fmt.Println("Error trying to load customer database: ", err.Error())
		}
	}
	if os.Getenv("LOAD_ADDRESS") == "S" {
		err = loadAddressTable()
		if err != nil {
			fmt.Println("Error trying to load customer database: ", err.Error())
		}
	}
	if os.Getenv("LOAD_PRODUCT_TYPE") == "S" {
		err = loadProductType()
		if err != nil {
			fmt.Println("Error trying to load customer database: ", err.Error())
		}
	}
	if os.Getenv("LOAD_PRODUCT") == "S" {
		err = loadProduct()
		if err != nil {
			fmt.Println("Error trying to load customer database: ", err.Error())
		}
	}
	if os.Getenv("LOAD_INVOICE") == "S" {
		err = loadInvoince()
		if err != nil {
			fmt.Println("Error trying to load customer database: ", err.Error())
		}
	}
	if os.Getenv("LOAD_ITEM") == "S" {
		err = loadItem()
		if err != nil {
			fmt.Println("Error trying to load customer database: ", err.Error())
		}
	}

}

func loadCustomersTable() error {
	fileContent, err := os.Open("./data/customer.json")

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var users Customers

	err = json.Unmarshal(byteResult, &users)

	if err != nil {
		log.Fatal(err)
		return err
	}

	for i := 0; i < len(users.CustomerList); i++ {
		fmt.Println("User Name: " + users.CustomerList[i].Nome)
		saveCustomer(users.CustomerList[i])
	}

	return nil
}

func saveCustomer(customer models.Cliente) (int, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	stmt, err := db.Prepare("INSERT INTO customer(name, cpf, phone_number) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	res, err := stmt.Exec(customer.Nome, customer.Telefone)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	fmt.Println(id)

	db.Close()

	return int(id), nil
}

func loadAddressTable() error {
	fileContent, err := os.Open("./data/address.json")

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var address Address

	err = json.Unmarshal(byteResult, &address)

	if err != nil {
		log.Fatal(err)
		return err
	}

	for i := 0; i < len(address.AddressList); i++ {
		SaveAddress(address.AddressList[i])
	}

	return nil
}

func SaveAddress(address models.Address) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
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

func loadProductType() error {
	fileContent, err := os.Open("./data/product_type.json")

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var productType ProductType

	err = json.Unmarshal(byteResult, &productType)

	if err != nil {
		log.Fatal(err)
		return err
	}

	for i := 0; i < len(productType.ProductTypeList); i++ {
		CreateProductType(productType.ProductTypeList[i])
	}

	return nil
}

func CreateProductType(product_type models.ProductType) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("insert into product_type (name, description) values (?,?)")
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

func loadProduct() error {
	fileContent, err := os.Open("./data/product.json")

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var product Product

	err = json.Unmarshal(byteResult, &product)

	if err != nil {
		log.Fatal(err)
		return err
	}

	for i := 0; i < len(product.ProductList); i++ {
		CreateProduct(product.ProductList[i])
	}

	return nil
}

func CreateProduct(product models.Product) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("insert into product (id_type, value_per_meter, total_value, thickness, color) values (?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(product.Id_tipo, product.Valor_metragem, product.Valor_total, product.Espessura, product.Cor)

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

func loadItem() error {
	fileContent, err := os.Open("./data/item.json")

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var item Item

	err = json.Unmarshal(byteResult, &item)

	if err != nil {
		log.Fatal(err)
		return err
	}

	for i := 0; i < len(item.ItemList); i++ {
		SaveItem(item.ItemList[i])
	}

	return nil
}

func SaveItem(item models.Item) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))

	stmt, err := db.Prepare("INSERT INTO item(id_invoice, id_product, amount, quantity, discount, metreage) values(?,?,?, ?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(item.Id_nota, item.Id_produto, item.Valor, item.Quantidade, item.Desconto, item.Metragem)

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

func loadInvoince() error {
	fileContent, err := os.Open("./data/invoice.json")

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("The File is opened successfully...")

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var invoice Invoice

	err = json.Unmarshal(byteResult, &invoice)

	if err != nil {
		log.Fatal(err)
		return err
	}

	for i := 0; i < len(invoice.InvoiceList); i++ {
		CreateNota(invoice.InvoiceList[i])
	}

	return nil
}

func CreateNota(invoice models.Nota) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("insert into invoice" +
		"(id_payment, id_customer, id_delivery_address, invoice_type, date, total_amount, total_discount)" +
		"values (?,?,?,?,?,?,?)")

	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(
		invoice.Id_pagamento,
		invoice.Id_cliente,
		invoice.Id_endereco_entrega,
		invoice.Tipo_nota,
		invoice.Data,
		invoice.Valor_total,
		invoice.Desconto_total)

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
