package api

import (
	"GoTaskManager/internal/app/api/rotas"
	"fmt"
	"time"
)

// Configurar: configura a API para utilização
func Configurar() {
	e := rotas.GerarEcho()

	e.Server.WriteTimeout = 30 * time.Second
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8000)))
}
