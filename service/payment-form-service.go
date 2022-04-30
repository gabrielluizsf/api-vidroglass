package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mariarobertap/api-vidroglass/interfaces"
	"github.com/mariarobertap/api-vidroglass/models"
	_ "github.com/mattn/go-sqlite3"
)

type paymentFormService struct {
	PaymentForm []models.PaymentForm
}

func NewPaymentFormService() interfaces.PaymentFormService {
	return &paymentFormService{}
}

func (c *paymentFormService) CreatePaymentForm(payment models.PaymentForm) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("insert into pagamento (forma_pagamento) values (?)")

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(payment.Paymentform)

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

func (c *paymentFormService) GetPaymentForm() ([]models.PaymentForm, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	rows, err := db.Query("SELECT * FROM pagamento")

	if err != nil {
		fmt.Println(err)
		return c.PaymentForm, err
	}

	c.PaymentForm = nil
	var id_pagamento int
	var forma_pagamento string

	for rows.Next() {
		err = rows.Scan(&id_pagamento, &forma_pagamento)
		if err != nil {
			fmt.Println(err)
			return c.PaymentForm, err
		}
		c.PaymentForm = append(c.PaymentForm, models.PaymentForm{id_pagamento, forma_pagamento})
	}

	rows.Close()

	return c.PaymentForm, nil

}

func (c *paymentFormService) GetPaymentFormByID(id_payment int) (models.PaymentForm, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	row := db.QueryRow("SELECT * FROM pagamento WHERE id_pagamento = ?", id_payment)

	var payment models.PaymentForm
	var id_pagamento int
	var forma_pagamento string
	err = row.Scan(&id_pagamento, &forma_pagamento)
	if err != nil {
		fmt.Println(err)
		return payment, err
	}

	payment = models.PaymentForm{id_pagamento, forma_pagamento}

	return payment, nil
}

func (c *paymentFormService) UpdatePaymentForm(payment models.PaymentForm) error {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("UPDATE pagamento SET forma_pagamento = ? WHERE id_pagamento = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(payment.Paymentform, payment.Id_payment)

	fmt.Println(res)
	if err != nil {
		return err
	}

	return nil
}
