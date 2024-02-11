package models

import (
	"errors"

	"gopkg.in/validator.v2"
)

// ValidarDados: valida se o json recebido está de acordo as validações
func ValidarDados(dados ...interface{}) (err error) {
	if len(dados) == 0 {
		return errors.New("dados não fornecidos")
	}

	for _, dado := range dados {
		if err = validator.Validate(dado); err != nil {
			return err
		}
	}
	return
}
