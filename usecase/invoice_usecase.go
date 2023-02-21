package usecase

import (
	"fmt"
	"vidroglass/gateway"
	"vidroglass/model"
	"vidroglass/repository"
)

type getInvoiceUseCase struct {
}

func NewInvoiceUseCase() gateway.InvoiceUseCase {
	return &getInvoiceUseCase{}
}

func (gc *getInvoiceUseCase) CreateInvoice(invoice model.Invoice) (model.Invoice, error) {

	invoice_id, err := repository.CreateInvoice(invoice)

	if err != nil {
		//log
		//todo return error

		return model.Invoice{}, nil

	}
	invoice.InvoiceID = invoice_id
	return invoice, nil

}

func (gc *getInvoiceUseCase) CreateInvoiceItem(item model.Item) (model.Item, error) {

	item_id, err := repository.CreateInvoiceItem(item)

	if err != nil {
		//log
		//todo return error

		return model.Item{}, nil

	}
	item.ItemID = item_id
	return item, nil

}

func (gc *getInvoiceUseCase) GetInvoices() ([]model.InvoiceObject, error) {

	invoices, err := repository.GetInvoices()

	if err != nil {
		//todo return	 error
		fmt.Println(err)
		return nil, nil

	}
	return invoices, nil
}
