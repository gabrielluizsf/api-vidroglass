package models

type Item struct {
	Id_item    int     `json:"id_item"`
	Id_produto int     `json:"id_produto,omitempty" binding:"required"`
	Id_nota    int     `json:"id_nota,omitempty" binding:"required"`
	Valor      float64 `json:"valor,omitempty"`
	Quantidade int     `json:"quantidade,omitempty" binding:"required"`
	Desconto   int     `json:"desconto,omitempty" `
	Metragem   float64 `json:"metragem,omitempty" binding:"required"`
}

type ItemPayload struct {
	Nome       string  `json:"nome,omitempty" binding:"required"`
	Cor        string  `json:"cor,omitempty" `
	Quantidade int     `json:"quantidade"`
	Valor      float64 `json:"valor,omitempty" binding:"required"`
	Desconto   float64 `json:"desconto,omitempty" binding:"required"`
	Metragem   int     `json:"metragem,omitempty"`
	Espessura  int     `json:"espessura,omitempty" binding:"required"`
}

type GoodResponseItem struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Item    Item   `json:"item"`
}
