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
	PaymentForm     models.PaymentForm
	PaymentFormList []models.PaymentForm
}

func NewPaymentFormService() interfaces.PaymentFormService {
	return &paymentFormService{}
}

func (c *paymentFormService) CreatePaymentForm(payment models.PaymentForm) (int, error) {

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("insert into payment (payment_form) values (?)")

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

	c.PaymentFormList = nil
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	rows, err := db.Query("SELECT * FROM payment")

	if err != nil {
		fmt.Println(err)
		return c.PaymentFormList, err
	}

	for rows.Next() {
		err = rows.Scan(
			&c.PaymentForm.Id_payment,
			&c.PaymentForm.Paymentform)

		if err != nil {
			fmt.Println(err)
			return c.PaymentFormList, err
		}
		c.PaymentFormList = append(c.PaymentFormList, c.PaymentForm)
	}

	rows.Close()

	return c.PaymentFormList, nil

}

func (c *paymentFormService) GetPaymentFormByID(id_payment int) (models.PaymentForm, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	row := db.QueryRow("SELECT * FROM payment WHERE id_pagamento = ?", id_payment)

	err = row.Scan(
		&c.PaymentForm.Id_payment,
		&c.PaymentForm.Paymentform)
	if err != nil {
		fmt.Println(err)
		return c.PaymentForm, err
	}
	return c.PaymentForm, nil
}

func (c *paymentFormService) UpdatePaymentForm(payment models.PaymentForm) error {
	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	defer db.Close()
	stmt, err := db.Prepare("UPDATE payment SET payment_form = ? WHERE id_pagamento = ?")

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
