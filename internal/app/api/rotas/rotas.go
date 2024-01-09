package rotas

import (
	"GoTaskManager/internal/app/api/middlewares"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// GerarEcho: retorna uma instância de Echo configurada inclusive as rotas
func GerarEcho() *echo.Echo {
	e := echo.New()

	configurarMiddlewares(e)
	configurarRotas(e)

	return e
}

// configurarMiddlewares: realiza a configuração dos middlewares
func configurarMiddlewares(e *echo.Echo) {
	e.Use(middleware.RequestID())

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:        true,
		LogStatus:     true,
		LogRequestID:  true,
		LogValuesFunc: middlewares.LoggerZeroLogPersonalizado,
	}))

	e.Use(middleware.CORS())
}

// configurarRotas: realiza a configuração das rotas
func configurarRotas(e *echo.Echo) {
	e.GET("/teste", func(c echo.Context) error {
		return c.String(http.StatusOK, time.Now().Format(time.RFC3339Nano))
	})

	RotasUsuario(e)
	RotasLogin(e)
}
