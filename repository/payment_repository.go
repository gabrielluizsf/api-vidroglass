package repository

import (
	"database/sql"
	"fmt"
	"vidroglass/model"

	_ "github.com/mattn/go-sqlite3"
)

func CreatePaymentMethod(payment model.PaymentMethod) (int, error) {

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	stmt, err := db.Prepare("insert into payment (payment_form) values (?)")

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(payment.PaymentMethod)

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

func GetPaymentsMethod() ([]model.PaymentMethod, error) {

	var paymentList []model.PaymentMethod
	var payment model.PaymentMethod

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	rows, err := db.Query("SELECT * FROM payment")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&payment.PaymentID,
			&payment.PaymentMethod)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		paymentList = append(paymentList, payment)
	}

	rows.Close()

	return paymentList, nil

}

func GetPaymentMethodByID(id_payment int) (model.PaymentMethod, error) {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	row := db.QueryRow("SELECT * FROM payment WHERE id_pagamento = ?", id_payment)
	var payment model.PaymentMethod

	err = row.Scan(
		&payment.PaymentID,
		&payment.PaymentMethod)
	if err != nil {
		fmt.Println(err)
		return model.PaymentMethod{}, err
	}
	return payment, nil
}

func UpdatePaymentMethod(payment model.PaymentMethod) error {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	stmt, err := db.Prepare("UPDATE payment SET payment_form = ? WHERE id_pagamento = ?")

	if err != nil {
		return err
	}

	res, err := stmt.Exec(payment.PaymentMethod, payment.PaymentID)

	fmt.Println(res)
	if err != nil {
		return err
	}

	return nil
}

func DeletePaymentByID(id_payment int) error {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM payment WHERE id_payment = ?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id_payment)

	fmt.Println(res)
	if err != nil {
		return err
	}

	return nil

}
