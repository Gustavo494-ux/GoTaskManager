package controllers

import (
	"github.com/labstack/echo/v4"
)

// ResponderErro: responde a requisição com uma mensagem de erro
func ResponderErro(c echo.Context, statusCode int, err error) error {
	resposta := map[string]interface{}{
		"erro": err.Error(),
		// "statuscode": statusCode,
	}
	return c.JSON(statusCode, resposta)
}

// ResponderString: responde a requisição com uma mensagem string
func ResponderString(c echo.Context, statusCode int, mensagem string) error {
	resposta := map[string]interface{}{
		"mensagem": mensagem,
		//"statuscode": statusCode,
	}
	return c.JSON(statusCode, resposta)
}
