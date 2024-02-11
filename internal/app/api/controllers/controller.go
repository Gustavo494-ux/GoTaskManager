package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// RetornarParametroInteiro: retorna um valor inteiro do parametro desejado
func RetornarParametroInteiro(c echo.Context, parametro string) (valor int) {
	valor, err := strconv.Atoi(c.Param(parametro))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	return
}
