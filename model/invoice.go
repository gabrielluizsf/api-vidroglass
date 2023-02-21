package model

type Invoice struct {
	InvoiceID     int     `json:"invoice_id"`
	PaymentID     int     `json:"payment_id"`
	CustomerID    int     `json:"customer_id"`
	InvoiceType   string  `json:"invoice_type" `
	Date          string  `json:"date" `
	TotalValue    float64 `json:"total_value" `
	TotalDiscount float64 `json:"total_discount" `
}

type InvoiceObject struct {
	InvoiceID     int          `json:"invoice_id"`
	PaymentMethod string       `json:"payment_method"`
	InvoiceType   string       `json:"invoice_type" `
	Date          string       `json:"date" `
	TotalValue    float64      `json:"total_value" `
	TotalDiscount float64      `json:"total_discount" `
	Item          []ItemObject `json:"item" `
	Customer      Customer     `json:"customer"`
}
