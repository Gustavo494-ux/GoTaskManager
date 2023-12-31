package middlewares

import (
	"GoTaskManager/pkg/pacotes/authentication"
	"GoTaskManager/pkg/pacotes/logger"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Authenticate: realiza a autenticação de todas as requisições que necessitam de autenticação
func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := validarToken(c); err != nil {
			return err
		}

		if err := authentication.IsExpirado(*c.Request()); err != nil {
			return err
		}

		if err := proximaFuncao(c, next); err != nil {
			return err
		}
		return nil
	}
}

// validarToken: verifica se o token é inválido, se for retorna um erro
func validarToken(e echo.Context) error {
	requisicao := *e.Request()
	if erro := authentication.ValidarToken(requisicao); erro != nil {
		logger.Logger().Info(fmt.Sprintf("Token %s inválido", authentication.ExtrairToken(requisicao)))
		return e.JSON(http.StatusNetworkAuthenticationRequired, "o token informado é inválido")
	}
	return nil
}

// proximaFuncao: executa a proxima função
func proximaFuncao(e echo.Context, next echo.HandlerFunc) error {
	requisicao := *e.Request()
	err := next(e)
	if err != nil {
		logger.Logger().Error(fmt.Sprintf("Ocorreu um erro na requisição, token: %s", authentication.ExtrairToken(requisicao)), err)
		return e.JSON(http.StatusUnauthorized, err)
	}
	return nil
}
