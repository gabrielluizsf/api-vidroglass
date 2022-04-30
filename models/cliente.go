package models

type Cliente struct {
	Id_cliente  int    `json:"idcliente"`
	Id_endereco int    `json:"idendereco,omitempty" binding:"required"`
	Nome        string `json:"nome,omitempty" binding:"required"`
	Cpf         string `json:"cpf,omitempty" binding:"required"`
	Telefone    string `json:"telefone,omitempty" binding:"required"`
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
