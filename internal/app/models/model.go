package models

import (
	"gopkg.in/validator.v2"
)

// ValidarDados: valida se o json recebido está de acordo as validações
func ValidarDados(v interface{}) (err error) {
	return validator.Validate(v)
}
