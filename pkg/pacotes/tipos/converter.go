package tipo

import (
	"GoTaskManager/pkg/pacotes/logger"
	"encoding/json"
	"fmt"
	"strconv"
)

type converter struct {
	valor string
}

// Converter: converte a variavel passada para json string
func Converter(valor any) converter {
	valorJson, err := json.Marshal(valor)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao converter a variavel valor para json", err, valor)
	}
	return converter{valor: string(valorJson)}
}

// Int: retorna o int
func (c converter) Int(base int, bitSize int) int64 {
	valorInteiro, err := strconv.ParseInt(c.valor, base, bitSize)
	if err != nil {
		logger.Logger().Error(fmt.Sprintf("Ocorreu um erro ao converter %s para int", c.valor), err, c.valor)
	}
	return valorInteiro
}

// String: retorna a string
func (c converter) String() string {
	return c.valor
}
