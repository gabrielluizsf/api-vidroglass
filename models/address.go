package models

type Address struct {
	Id_address  int    `json:"id_address"`
	Id_customer int    `json:"id_customer,omitempty" validate:"required"`
	Street      string `json:"street,omitempty" validate:"required"`
	Number      int64  `json:"number" binding:"required"`
	Cep         string `json:"cep,omitempty" validate:"required"`
	City        string `json:"city,omitempty" validate:"required"`
	State       string `json:"state,omitempty" validate:"required"`
}

type GoodResponseAddress struct {
	Message string  `json:"message"`
	Status  string  `json:"status"`
	Address Address `json:"Address"`
}
