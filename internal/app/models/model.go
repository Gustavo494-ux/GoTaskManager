package models

import (
	"errors"

	"gopkg.in/validator.v2"
)

// ValidarDados: valida se o json recebido está de acordo as validações
func ValidarDados(v interface{}) (err error) {
	if v == nil {
		return errors.New("o corpo da requisição não pode ser nulo")
	}
	return validator.Validate(v)
}
