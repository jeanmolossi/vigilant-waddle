package adapter

import "errors"

type AuthCredentials struct {
	Email    string `json:"email" validate:"required,email" example:"john@doe.com"`
	Password string `json:"password" validate:"required,min=6,max=64" example:"123456"`
}

func (a *AuthCredentials) GetErrorMap() map[string]map[string]error {
	return map[string]map[string]error{
		"email": {
			"required": errors.New("email is required"),
			"email":    errors.New("email is invalid"),
		},
		"password": {
			"required": errors.New("password is required"),
			"min":      errors.New("password must be at least 6 characters"),
			"max":      errors.New("password must be at most 64 characters"),
		},
	}
}
