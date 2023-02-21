package repository

import (
	"database/sql"
	"fmt"
	"time"
	"vidroglass/model"

	_ "github.com/mattn/go-sqlite3"
)

func CreateInvoice(invoice model.Invoice) (int, error) {

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	stmt, err := db.Prepare("insert into invoice" +
		"(id_payment, id_customer, invoice_type, date)" +
		"values (?,?,?,?)")

	if err != nil {
		return 0, err
	}
	invoice.Date = fmt.Sprintf("%v", time.Now())

	res, err := stmt.Exec(
		invoice.PaymentID,
		invoice.CustomerID,
		invoice.InvoiceType,
		invoice.Date)

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

func GetInvoices() ([]model.InvoiceObject, error) {

	var invoice model.InvoiceObject
	var invoiceList []model.InvoiceObject

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()
	rows, err := db.Query("select " +
		"i.id_invoice, " +
		"p.payment_form, " +
		"c.name, " +
		"c.state, " +
		"c.city, " +
		"c.street, " +
		"c.number, " +
		"i.invoice_type, " +
		"i.date, " +
		"i.total_amount, " +
		"i.total_discount " +
		"from " +
		"invoice i " +
		"JOIN " +
		"customer c on c.id_customer = i.id_customer " +
		"JOIN " +
		"payment p on p.id_payment = i.id_payment ")

	if err != nil {
		//log
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&invoice.InvoiceID,
			&invoice.PaymentMethod,
			&invoice.Customer.Nome,
			&invoice.Customer.Address.State,
			&invoice.Customer.Address.City,
			&invoice.Customer.Address.Street,
			&invoice.Customer.Address.Number,
			&invoice.InvoiceType,
			&invoice.Date,
			&invoice.TotalValue,
			&invoice.TotalDiscount)

		if err != nil {
			return nil, err
		}

		item, err := getInvoiceItems(invoice.InvoiceID)
		if err != nil {
			continue
		}

		fmt.Println(invoice.InvoiceID)
		fmt.Println(item)
		totalValue, _ := getInvoiceTotalValue(invoice.InvoiceID)

		invoice.TotalValue = totalValue
		invoice.Item = item
		fmt.Println(invoice.Customer.Address.State)

		invoiceList = append(invoiceList, invoice)
		item = nil
		totalValue = 0
	}

	rows.Close()

	return invoiceList, nil
}

func getInvoiceItems(invoice_id int) ([]model.ItemObject, error) {
	var itemObject model.ItemObject
	var itemObjects []model.ItemObject

	db, err := sql.Open("sqlite3", "database.db")
	rows, err := db.Query("SELECT "+
		"i.quantity, i.value, i.discount,"+
		"i.metreage, p.color, tp.name, p.value_per_meter "+
		"FROM item as i "+
		"join product as p on p.id_product = i.id_product "+
		"join product_type as tp on tp.id_type = p.id_type "+
		"where i.id_invoice = ?", invoice_id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&itemObject.Quantity,
			&itemObject.Value,
			&itemObject.Discount,
			&itemObject.Metragem,
			&itemObject.Color,
			&itemObject.ProductName,
			&itemObject.ValuePerMeter)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		itemObjects = append(itemObjects, itemObject)
	}

	rows.Close()
	db.Close()

	return itemObjects, nil
}

func getInvoiceTotalValue(invoice_id int) (float64, error) {
	var totalValue float64

	db, err := sql.Open("sqlite3", "database.db")
	row := db.QueryRow("select sum(value) from item where id_invoice = ?", invoice_id)

	err = row.Scan(
		&totalValue)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return totalValue, nil
}

func updateInvoiceTotalValue(invoice_id int) (int, error) {

	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	stmt, err := db.Prepare("UPDATE " +
		"invoice " +
		"set " +
		"total_amount = (select sum(value) from item where id_invoice = ?) " +
		"where " +
		"id_invoice = ?")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	res, err := stmt.Exec(invoice_id, invoice_id)

	if err != nil {
		fmt.Println(err)

		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)

		return 0, err
	}

	fmt.Println(id)

	return int(id), nil

}

func CreateInvoiceItem(item model.Item) (int, error) {
	db, err := sql.Open("sqlite3", "database.db")
	defer db.Close()

	stmt, err := db.Prepare("insert into item" +
		"(id_invoice, id_product, value, quantity, discount, metreage)" +
		"values (?,?,?,?,?,?)")

	if err != nil {
		return 0, err
	}
	item.Value = getItemValue(item)

	res, err := stmt.Exec(
		item.InvoiceID,
		item.ProductID,
		item.Value,
		item.Quantity,
		item.Discount,
		item.Metragem)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	_, err = updateInvoiceTotalValue(item.InvoiceID)
	if err != nil {
		return 0, err
	}

	fmt.Println(id)

	return int(id), nil
}

func getItemValue(invoiceItem model.Item) float64 {

	product, err := getProductValue(invoiceItem.ProductID)

	fmt.Println(product)
	if err != nil {
		return 0
	}

	if product.TotalValue != 0 {
		itemValue := product.TotalValue * float64(invoiceItem.Quantity)
		return itemValue
	} else if product.ValuePerMeter != 0 {
		fmt.Println("Aqui")

		itemValue := (product.ValuePerMeter * invoiceItem.Metragem) * float64(invoiceItem.Quantity)
		return itemValue
	}

	return 0
}

// func GetInvoiceByID(id_invoice int) (model.InvoicePayload, error) {

// 	var invoice model.Invoice
// 	var invoicePayload model.InvoicePayload
// 	fmt.Println(id_invoice)

// 	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
// 	defer db.Close()
// 	row := db.QueryRow("SELECT * FROM invoice WHERE id_invoice = ?", id_invoice)

// 	err = row.Scan(&invoice.Id_invoice,
// 		&invoice.Id_pagamento,
// 		&invoice.Id_cliente,
// 		&invoice.Id_endereco_entrega,
// 		&invoice.Tipo_invoice,
// 		&invoice.Data,
// 		&invoice.Valor_total,
// 		&invoice.Desconto_total)

// 	if err != nil {
// 		fmt.Println(err)
// 		return invoicePayload, err
// 	}

// 	invoicePayload, err = c.buildInvoicePayload(invoice)

// 	if err != nil {
// 		fmt.Println(err)
// 		return invoicePayload, err
// 	}

// 	return invoicePayload, nil
// }

// func UpdateInvoice(invoice model.Invoice) (model.Invoice, error) {
// 	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
// 	defer db.Close()
// 	stmt, err := db.Prepare("UPDATE invoice SET id_payment = ?, id_customer = ?,  id_delivery_address = ?, invoice_type = ?,  date = ?, total_amount = ?, total_discount = ? WHERE id_invoice = ?")

// 	if err != nil {
// 		return invoice, err
// 	}

// 	total_value, err := c.getTotalValueInvoice(invoice.Id_invoice)
// 	fmt.Println(total_value)

// 	if err != nil {
// 		return invoice, err
// 	}
// 	res, err := stmt.Exec(
// 		invoice.Id_pagamento,
// 		invoice.Id_cliente,
// 		invoice.Id_endereco_entrega,
// 		invoice.Tipo_invoice,
// 		invoice.Data,
// 		total_value,
// 		invoice.Desconto_total,
// 		invoice.Id_invoice)

// 	fmt.Println(res)
// 	if err != nil {
// 		return invoice, err
// 	}

// 	invoice.Valor_total = total_value

// 	return invoice, nil
// }

// func getTotalValueInvoice(id_invoice int) (float64, error) {

// 	var totalValue float64

// 	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
// 	defer db.Close()
// 	row := db.QueryRow("SELECT sum(amount*quantity) FROM item where id_invoice = ?", id_invoice)

// 	if err != nil {
// 		return 0, err
// 	}

// 	err = row.Scan(&totalValue)
// 	if err != nil {
// 		fmt.Println(err)
// 		return 0, err
// 	}

// 	return totalValue, nil
// }

// func buildInvoicePayload(invoice model.Invoice) (model.InvoicePayload, error) {
// 	var clienteService interfaces.ClienteService = NewClienteService()
// 	var ItemService interfaces.ItemService = NewItemService()
// 	var AddressService interfaces.AddressService = NewAddressService()

// 	cliente, err := clienteService.GetClientById(invoice.Id_cliente)
// 	if err != nil {
// 		fmt.Println("Cliente não econtrado")
// 		return model.InvoicePayload{}, err
// 	}
// 	Item, err := ItemService.GetInvoiceItens(invoice.Id_invoice)
// 	if err != nil {
// 		fmt.Println("Itens não encontrados")

// 		return model.InvoicePayload{}, err
// 	}
// 	endereco, err := AddressService.GetAddressByID(invoice.Id_endereco_entrega)
// 	if err != nil {
// 		fmt.Println("Endereço não encontrado")

// 		return model.InvoicePayload{}, err
// 	}
// 	cliente.Endereco = endereco

// 	invoicePayload := model.InvoicePayload{
// 		Invoice: model.InvoiceDetails{
// 			Id_invoice:     invoice.Id_invoice,
// 			Id_pagamento:   invoice.Id_pagamento,
// 			Tipo_invoice:   invoice.Tipo_invoice,
// 			Data:           invoice.Data,
// 			Valor_total:    c.getInvoiceTotalValue(Item),
// 			Desconto_total: invoice.Desconto_total,
// 			Item:           Item,
// 			Cliente:        cliente,
// 		},
// 	}

// 	fmt.Printf("Passando aqui")
// 	return invoicePayload, nil
// }

// func getInvoiceTotalValue(items []model.ItemPayload) float64 {
// 	var total float64
// 	for i := 0; i < len(items); i++ {
// 		total += items[i].Valor
// 	}

// 	return total
// }
