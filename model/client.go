package model

type Customer struct {
	ID       int     `json:"idcliente"`
	Nome     string  `json:"nome,omitempty" binding:"required"`
	Telefone string  `json:"telefone,omitempty"`
	Address  Address `json:"address,omitempty"`
}

type Address struct {
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
	Number  string `json:"number,omitempty"`
	ZipCode string `json:"zip_code,omitempty"`
}
