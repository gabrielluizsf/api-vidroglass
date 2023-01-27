package models

type Cliente struct {
	Id_cliente int     `json:"idcliente"`
	Nome       string  `json:"nome,omitempty" binding:"required"`
	Telefone   string  `json:"telefone,omitempty" binding:"required"`
	Endereco   Address `json:"address,omitempty"`
}

type GoodResponse struct {
	Message string  `json:"message"`
	Status  string  `json:"status"`
	Cliente Cliente `json:"cliente"`
}

type BadResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Erro    string `json:"erro"`
}
