package models

type Product struct {
	Id_produto     int     `json:"id_produto"`
	Id_tipo        int     `json:"id_tipo"`
	Valor_metragem float32 `json:"valor_metragem,omitempty" binding:"required"`
	Espessura      float32 `json:"espessura" `
	Cor            string  `json:"cor" binding:"required"`
}

type GoodResponseProduct struct {
	Message string  `json:"message"`
	Status  string  `json:"status"`
	Product Product `json:"payment_form"`
}
