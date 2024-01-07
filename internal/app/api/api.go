package api

import (
	"GoTaskManager/internal/app/api/rotas"
	"GoTaskManager/pkg/pacotes/logger"
	"fmt"
	"strconv"
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
	var err error
	e := rotas.GerarEcho()

	PortaApi, err = strconv.Atoi(porta)
	if err != nil {
		PortaApi = PortaApiPadrao
	}

	StartarApi(e, time.Second*30)
}

// StartarApi: realiza o start da api
func StartarApi(e *echo.Echo, timeout time.Duration) {
	e.Server.WriteTimeout = timeout
	if err := e.Start(fmt.Sprintf(":%s", strconv.Itoa(PortaApi))); err != nil {
		logger.Logger().Fatal("Ocorreu um erro ao startar a api", err)
	}
}
