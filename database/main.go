package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func init() {

	//os.Create("data.db")

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/*
	   _, err = db.Exec("CREATE TABLE `customers` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `first_name` VARCHAR(255) NOT NULL, `last_name` VARCHAR(255) NOT NULL)")
	   if err != nil {
	       fmt.Println(err)
	       os.Exit(1)
	   }
	*/

	getAll(db)
	db.Close()

}

func getAll(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM endereco")
	checkErr(err)
	var uid int
	var rua string
	var numero int
	var cep string
	var cidade string
	var estado string

	for rows.Next() {
		err = rows.Scan(&uid, &rua, &numero, &cep, &cidade, &estado)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(rua)
		fmt.Println(numero)
		fmt.Println(cep)
		fmt.Println(cidade)
		fmt.Println(estado)

	}

	rows.Close()
}

func createCustomer(db *sql.DB) {
	// insert

	stmt, err := db.Prepare("INSERT INTO customers(first_name, last_name) values(?,?)")
	checkErr(err)

	res, err := stmt.Exec("Maria", "Roberta")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
