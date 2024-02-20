package middlewares

import (
	"net/http"

	"github.com/Gustavo494-ux/PacotesGolang/configuracoes"
	"github.com/Gustavo494-ux/PacotesGolang/logger"
	"github.com/labstack/echo/v4"
)

// Authenticate: realiza a autenticação de todas as requisições que necessitam de autenticação
func BancoDeDadosDisponivel(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bancoIndisponivel := func(err error) error {
			logger.Logger().Error("banco de dados indisponível", err)
			return c.JSON(http.StatusServiceUnavailable, "banco de dados indisponivel")
		}

		db, err := configuracoes.BancoPrincipalGORM.DB()
		if err != nil {
			return bancoIndisponivel(err)
		}

		if err = db.Ping(); err != nil {
			return bancoIndisponivel(err)
		}

		return ProximaFuncao(c, next)
	}
}
