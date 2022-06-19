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
	NotaList []models.Nota
}

func NewNotaService() interfaces.NotaService {
	return &notaService{}
}

func (c *notaService) CreateNota() (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("insert into nota" +
		"(id_pagamento)" +
		"values (0)")

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
	res, err := stmt.Exec()

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

func (c *notaService) GetNota() ([]models.Nota, error) {

	var nota models.Nota
	c.NotaList = nil

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	rows, err := db.Query("SELECT * FROM nota")

	if err != nil {
		fmt.Println(err)
		return c.NotaList, err
	}

	for rows.Next() {
		err = rows.Scan(&nota.Id_nota, &nota.Id_pagamento,
			&nota.Id_cliente, &nota.Id_endereco_entrega,
			&nota.Tipo_nota, &nota.Data, &nota.Valor_total,
			&nota.Desconto_total)

		fmt.Println(nota.Id_pagamento)

		if err != nil {
			fmt.Println(err)
			return c.NotaList, err
		}
		c.NotaList = append(c.NotaList, nota)
	}

	rows.Close()

	return c.NotaList, nil

}

func (c *notaService) GetNotaByID(id_nota int) (models.NotaPayload, error) {

	var nota models.Nota
	var notaPayload models.NotaPayload

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	row := db.QueryRow("SELECT * FROM nota WHERE id_nota = ?", id_nota)

	err = row.Scan(&nota.Id_nota, &nota.Id_pagamento,
		&nota.Id_cliente, &nota.Id_endereco_entrega,
		&nota.Tipo_nota, &nota.Data, &nota.Valor_total,
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
	stmt, err := db.Prepare("UPDATE nota SET id_pagamento = ?, id_cliente = ?,  id_endereco_entrega = ?, tipo_nota = ?,  valor_total = ? WHERE id_nota = ?")

	if err != nil {
		return nota, err
	}

	total_value, err := c.getTotalValueInvoice(nota.Id_nota)
	fmt.Println(total_value)

	if err != nil {
		return nota, err
	}
	res, err := stmt.Exec(nota.Id_pagamento,
		nota.Id_cliente,
		nota.Id_endereco_entrega,
		nota.Tipo_nota, total_value, nota.Id_nota)

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
	row := db.QueryRow("SELECT sum(valor*quantidade) FROM item where id_nota = ?", id_nota)

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
	var notaPayload models.NotaPayload
	var clienteService interfaces.ClienteService = NewClienteService()
	var ItemService interfaces.ItemService = NewItemService()
	var AddressService interfaces.AddressService = NewAddressService()

	cliente, err := clienteService.GetClientById(nota.Id_cliente)
	if err != nil {
		return notaPayload, err
	}
	Item, err := ItemService.GetInvoiceItens(nota.Id_nota)
	if err != nil {
		return notaPayload, err
	}
	endereco, err := AddressService.GetAddressByID(nota.Id_endereco_entrega)
	if err != nil {
		return notaPayload, err
	}

	notaPayload = models.NotaPayload{
		Id_nota:        nota.Id_nota,
		Id_pagamento:   nota.Id_pagamento,
		Tipo_nota:      nota.Tipo_nota,
		Data:           nota.Data,
		Valor_total:    nota.Valor_total,
		Desconto_total: nota.Desconto_total,
		Item:           Item,
		Cliente:        cliente,
		Address:        endereco}

	return notaPayload, nil
}
