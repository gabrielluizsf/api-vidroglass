package models

type ProductType struct {
	Id_tipo   int    `json:"id_tipo"`
	Nome      string `json:"nome,omitempty" binding:"required"`
	Descricao string `json:"descricao"`
}

type GoodResponseProductType struct {
	Message     string      `json:"message"`
	Status      string      `json:"status"`
	ProductType ProductType `json:"payment_form"`
}
