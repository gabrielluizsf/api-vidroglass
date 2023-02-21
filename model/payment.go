package model

type PaymentMethod struct {
	PaymentID     int    `json:"idpagamento"`
	PaymentMethod string `json:"pagamento,omitempty" binding:"required"`
}
