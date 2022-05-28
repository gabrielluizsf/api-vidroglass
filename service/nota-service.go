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
	Nota []models.Nota
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

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	rows, err := db.Query("SELECT * FROM nota")

	if err != nil {
		fmt.Println(err)
		return c.Nota, err
	}
	var id_nota int
	var id_pagamento int
	var id_cliente int
	var id_endereco int
	var tipo_nota string
	var data string
	var valor_total float64
	var desconto_total float64
	c.Nota = nil

	for rows.Next() {
		err = rows.Scan(&id_nota, &id_pagamento, &id_cliente, &id_endereco, &tipo_nota, &data, &valor_total, &desconto_total)
		fmt.Println(id_pagamento)
		if err != nil {
			fmt.Println(err)
			return c.Nota, err
		}
		c.Nota = append(c.Nota, models.Nota{id_nota, id_pagamento, id_cliente, id_endereco, tipo_nota, data, valor_total, desconto_total})
	}

	rows.Close()

	return c.Nota, nil

}

func (c *notaService) GetNotaByID(id_nota int) (models.Nota, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	row := db.QueryRow("SELECT * FROM nota WHERE id_nota = ?", id_nota)

	var nota models.Nota
	var id_pagamento int
	var id_cliente int
	var id_endereco int
	var tipo_nota string
	var data string
	var valor_total float64
	var desconto_total float64

	err = row.Scan(&id_nota, &id_pagamento, &id_cliente, &id_endereco, &tipo_nota, &data, &valor_total, &desconto_total)
	if err != nil {
		fmt.Println(err)
		return nota, err
	}

	nota = models.Nota{id_nota, id_pagamento, id_cliente, id_endereco, tipo_nota, data, valor_total, desconto_total}

	return nota, nil
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
