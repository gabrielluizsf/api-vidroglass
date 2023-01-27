package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
	_ "github.com/mattn/go-sqlite3"
)

type notaService struct {
	Nota     models.Nota
	NotaList []models.NotaPayload
}

func NewNotaService() interfaces.NotaService {
	return &notaService{}
}

func (c *notaService) CreateNota(nota models.Nota) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("insert into invoice" +
		"(id_payment, id_customer, id_delivery_address, invoice_type, date)" +
		"values (?,?,?,?,?)")

	if err != nil {
		return 0, err
	}
	/*
		total_value, err := c.getTotalValueInvoice(nota.Id_nota)

		if err != nil {
			return 0, err
		}

		res, err := stmt.Exec(nota.Id_pagamento, nota.Id_cliente, nota.Id_endereco_entrega, nota.Tipo_nota, nota.Data, total_value, 0)

		if err != nil {
			return 0, err
		}
	*/
	res, err := stmt.Exec(nota.Id_pagamento, nota.Id_cliente, nota.Id_endereco_entrega, nota.Tipo_nota, nota.Data)

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

func (c *notaService) GetNota() ([]models.NotaPayload, error) {

	var nota models.Nota
	c.NotaList = nil

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	rows, err := db.Query("SELECT * FROM invoice")

	if err != nil {
		fmt.Println(err)
		return c.NotaList, err
	}

	for rows.Next() {
		err = rows.Scan(
			&nota.Id_nota,
			&nota.Id_pagamento,
			&nota.Id_cliente,
			&nota.Id_endereco_entrega,
			&nota.Tipo_nota,
			&nota.Data,
			&nota.Valor_total,
			&nota.Desconto_total)

		fmt.Println(nota.Id_pagamento)

		if err != nil {
			fmt.Println(err)
			return c.NotaList, err
		}

		notaPayload, err := c.buildNotaPayload(nota)

		if err != nil {
			fmt.Println(err)
			continue
		}

		c.NotaList = append(c.NotaList, notaPayload)
	}

	rows.Close()

	return c.NotaList, nil

}

func (c *notaService) GetNotaByID(id_nota int) (models.NotaPayload, error) {

	var nota models.Nota
	var notaPayload models.NotaPayload
	fmt.Println(id_nota)

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	row := db.QueryRow("SELECT * FROM invoice WHERE id_invoice = ?", id_nota)

	err = row.Scan(&nota.Id_nota,
		&nota.Id_pagamento,
		&nota.Id_cliente,
		&nota.Id_endereco_entrega,
		&nota.Tipo_nota,
		&nota.Data,
		&nota.Valor_total,
		&nota.Desconto_total)

	if err != nil {
		fmt.Println(err)
		return notaPayload, err
	}

	notaPayload, err = c.buildNotaPayload(nota)

	if err != nil {
		fmt.Println(err)
		return notaPayload, err
	}

	return notaPayload, nil
}

func (c *notaService) UpdateNota(nota models.Nota) (models.Nota, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("UPDATE invoice SET id_payment = ?, id_customer = ?,  id_delivery_address = ?, invoice_type = ?,  date = ?, total_amount = ?, total_discount = ? WHERE id_invoice = ?")

	if err != nil {
		return nota, err
	}

	total_value, err := c.getTotalValueInvoice(nota.Id_nota)
	fmt.Println(total_value)

	if err != nil {
		return nota, err
	}
	res, err := stmt.Exec(
		nota.Id_pagamento,
		nota.Id_cliente,
		nota.Id_endereco_entrega,
		nota.Tipo_nota,
		nota.Data,
		total_value,
		nota.Desconto_total,
		nota.Id_nota)

	fmt.Println(res)
	if err != nil {
		return nota, err
	}

	nota.Valor_total = total_value

	return nota, nil
}

func (c *notaService) getTotalValueInvoice(id_nota int) (float64, error) {

	var totalValue float64

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	row := db.QueryRow("SELECT sum(amount*quantity) FROM item where id_invoice = ?", id_nota)

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

func (c *notaService) buildNotaPayload(nota models.Nota) (models.NotaPayload, error) {
	var clienteService interfaces.ClienteService = NewClienteService()
	var ItemService interfaces.ItemService = NewItemService()
	var AddressService interfaces.AddressService = NewAddressService()

	cliente, err := clienteService.GetClientById(nota.Id_cliente)
	if err != nil {
		fmt.Println("Cliente não econtrado")
		return models.NotaPayload{}, err
	}
	Item, err := ItemService.GetInvoiceItens(nota.Id_nota)
	if err != nil {
		fmt.Println("Itens não encontrados")

		return models.NotaPayload{}, err
	}
	endereco, err := AddressService.GetAddressByID(nota.Id_endereco_entrega)
	if err != nil {
		fmt.Println("Endereço não encontrado")

		return models.NotaPayload{}, err
	}
	cliente.Endereco = endereco

	notaPayload := models.NotaPayload{
		Nota: models.NotaDetails{
			Id_nota:        nota.Id_nota,
			Id_pagamento:   nota.Id_pagamento,
			Tipo_nota:      nota.Tipo_nota,
			Data:           nota.Data,
			Valor_total:    c.getInvoiceTotalValue(Item),
			Desconto_total: nota.Desconto_total,
			Item:           Item,
			Cliente:        cliente,
		},
	}

	fmt.Printf("Passando aqui")
	return notaPayload, nil
}

func (c *notaService) getInvoiceTotalValue(items []models.ItemPayload) float64 {
	var total float64
	for i := 0; i < len(items); i++ {
		total += items[i].Valor
	}

	return total
}
