package api

import (
	"GoTaskManager/internal/app/api/rotas"
	"GoTaskManager/pkg/pacotes/logger"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	PortaApiPadrao = 8000
)

var (
	PortaApi int
)

// Configurar: configura a API para utilização
func ConfigurarApi(porta string) {
	e := rotas.GerarEcho()

	StartarApi(e, porta, time.Second*30)
}

// StartarApi: realiza o start da api
func StartarApi(e *echo.Echo, porta string, timeout time.Duration) {
	e.Server.WriteTimeout = timeout
	if err := e.Start(fmt.Sprintf(":%s", porta)); err != nil {
		logger.Logger().Fatal("Ocorreu um erro ao startar a api", err)
	}
}
