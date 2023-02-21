package util

import "vidroglass/model"

func BuildErrorResponse(message string, description string) model.Reponse {

	return model.Reponse{
		Type:    "error",
		Success: false,
		Detail: model.ErrorReponse{
			ErrorMessage: message,
			Description:  description,
		},
	}
}

func BuildSuccessReponse(message string) model.Reponse {

	return model.Reponse{
		Type:    "error",
		Success: false,
		Detail: model.SuccessMessage{
			ErrorMessage: message,
		},
	}
}
