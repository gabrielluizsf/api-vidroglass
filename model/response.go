package model

type Reponse struct {
	Type    string
	Success bool
	Detail  interface{}
}

type ErrorReponse struct {
	ErrorMessage string
	Description  string
}

type SuccessMessage struct {
	ErrorMessage string
}
