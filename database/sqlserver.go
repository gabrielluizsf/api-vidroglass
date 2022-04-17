package  main

import (
    "database/sql"
    "fmt"
    _ "github.com/denisenkom/go-mssqldb"
)

func main() {


	// Create connection string
    connString := fmt.Sprintf("server=DESKTOP-IG3M3Q8\\SQLEXPRESS;user id=;password=;port=51287;encrypt=disable")
    db, err := sql.Open("mssql", connString)
    if err != nil {
        fmt.Println("Cannot connect: ", err.Error())
        return
    }
    err = db.Ping()
    if err != nil {
        fmt.Println("Cannot connect: ", err.Error())
        return
    }
    defer db.Close()
}