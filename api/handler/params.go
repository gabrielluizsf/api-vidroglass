package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getCustomerID(ctx *gin.Context) (int, error) {
	customer_id := ctx.Param("customerID")

	if (customer_id) == "" {
		return 0, fmt.Errorf("Invalid customer id")
	}

	id, err := strconv.Atoi(customer_id)

	if err != nil {
		return 0, fmt.Errorf("Invalid customer id")
	}

	return id, nil

}

func getInvoiceID(ctx *gin.Context) (int, error) {
	invoice_id := ctx.Param("invoiceID")

	if (invoice_id) == "" {
		return 0, fmt.Errorf("Invalid invoice id")
	}

	id, err := strconv.Atoi(invoice_id)

	if err != nil {
		return 0, fmt.Errorf("Invalid invoice id")
	}

	return id, nil

}

func getProductID(ctx *gin.Context) (int, error) {
	product_id := ctx.Param("productID")

	if (product_id) == "" {
		return 0, fmt.Errorf("Invalid product id")
	}

	id, err := strconv.Atoi(product_id)

	if err != nil {
		return 0, fmt.Errorf("Invalid product id")
	}

	return id, nil

}

func getPaymentID(ctx *gin.Context) (int, error) {
	payment_id := ctx.Param("paymentID")

	if (payment_id) == "" {
		return 0, fmt.Errorf("Invalid payment id")
	}

	id, err := strconv.Atoi(payment_id)

	if err != nil {
		return 0, fmt.Errorf("Invalid payment id")
	}

	return id, nil

}
