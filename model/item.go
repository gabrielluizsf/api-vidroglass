package model

type Item struct {
	ItemID    int     `json:"id_item"`
	InvoiceID int     `json:"id_nota,omitempty" binding:"required"`
	ProductID int     `json:"id_produto,omitempty" binding:"required"`
	Value     float64 `json:"valor,omitempty"`
	Quantity  int     `json:"quantidade,omitempty" binding:"required"`
	Discount  float64 `json:"desconto,omitempty" `
	Metragem  float64 `json:"metragem,omitempty" binding:"required"`
}

type ItemObject struct {
	ItemID        int     `json:"id_item"`
	ProductName   string  `json:"name,omitempty" binding:"required"`
	Value         float64 `json:"valor,omitempty"`
	Quantity      int     `json:"quantidade,omitempty" binding:"required"`
	Discount      float64 `json:"desconto,omitempty" `
	Color         string  `json:"color,omitempty" `
	ValuePerMeter float64 `json:"value_per_meter,omitempty" `
	Metragem      float64 `json:"metragem,omitempty" binding:"required"`
}
