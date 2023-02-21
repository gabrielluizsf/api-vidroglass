package gateway

import "vidroglass/model"

type InvoiceUseCase interface {
	GetInvoices() ([]model.InvoiceObject, error)
	CreateInvoice(model.Invoice) (model.Invoice, error)
	CreateInvoiceItem(model.Item) (model.Item, error)
}
