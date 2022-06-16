package models

type PaymentForm struct {
	Id_payment  int    `json:"idpagamento"`
	Paymentform string `json:"pagamento,omitempty" binding:"required"`
}
//teste
type GoodResponsePayment struct {
	Message     string      `json:"message"`
	Status      string      `json:"status"`
	PaymentForm PaymentForm `json:"payment_form"`
}
