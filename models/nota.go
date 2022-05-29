package models

type Nota struct {
	Id_nota             int     `json:"id_nota"`
	Id_pagamento        int     `json:"Id_pagamento"`
	Id_cliente          int     `json:"Id_cliente"`
	Id_endereco_entrega int     `json:"Id_endereco_entrega"`
	Tipo_nota           string  `json:"tipo_nota" `
	Data                string  `json:"data" `
	Valor_total         float64 `json:"valor_total" `
	Desconto_total      float64 `json:"desconto_total" `
}

type NotaPayload struct {
	Id_nota        int           `json:"id_nota"`
	Id_pagamento   int           `json:"Id_pagamento"`
	Tipo_nota      string        `json:"tipo_nota" `
	Data           string        `json:"data" `
	Valor_total    float64       `json:"valor_total" `
	Desconto_total float64       `json:"desconto_total" `
	Item           []ItemPayload `json:"itens" `
	Cliente        Cliente       `json:"cliente"`
	Address        Address       `json:"endereco"`
}

type GoodResponseNota struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Id_nota int    `json:"id_nota"`
}
type GoodResponseNotaObjetc struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Nota    Nota   `json:"nota"`
}

type BadResponseNota struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Erro    string `json:"erro"`
}
