package models

type Cliente struct {
	Id_cliente  int    `json:"idcliente"`
	Id_endereco int    `json:"idendereco"`
	Nome        string `json:"nome"`
	Cpf         string `json:"cpf"`
	Telefone    string `json:"telefone"`
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
