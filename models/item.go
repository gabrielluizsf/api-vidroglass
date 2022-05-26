package models

type Item struct {
	Id_item    int     `json:"id_item"`
	Id_produto int     `json:"id_produto,omitempty" binding:"required"`
	Id_nota    string  `json:"id_nota,omitempty" binding:"required"`
	Valor      float64 `json:"valor,omitempty" binding:"required"`
	Quantidade int     `json:"quantidade,omitempty" binding:"required"`
	Desconto   int     `json:"desconto,omitempty" binding:"required"`
	Metragem   float64 `json:"metragem,omitempty" binding:"required"`
}

type GoodResponseItem struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Item    Item   `json:"item"`
}
