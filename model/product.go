package model

type Product struct {
	ProductID     int     `json:"id_produto"`
	ProductTypeID int     `json:"id_tipo"`
	ValuePerMeter float64 `json:"valor_metragem,omitempty"`
	TotalValue    float64 `json:"valor_total,omitempty" `
	Espessura     float64 `json:"espessura" `
	Color         string  `json:"cor"`
}

type ProductObject struct {
	ProductID     int     `json:"id_produto"`
	ProductType   string  `json:"type"`
	ValuePerMeter float64 `json:"valor_metragem,omitempty"`
	TotalValue    float64 `json:"valor_total,omitempty" `
	Espessura     float64 `json:"espessura" `
	Color         string  `json:"cor"`
}

type ProductType struct {
	TypeID      int    `json:"id_tipo"`
	Name        string `json:"nome,omitempty" binding:"required"`
	Description string `json:"descricao"`
}
