package configuracoes

import (
	"GoTaskManager/internal/app/api/rotas"
	"GoTaskManager/pkg/pacotes/logger"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	PortaApiPadrao = 8000
)

var (
	PortaApi int
)

// Configurar: configura a API para utilização
func ConfigurarApi(CaminhoRelativoArquivoConfiguracao string) {
	caminhoArquivoConfiguracao := PrepararCaminhoArquivo(CaminhoRelativoArquivoConfiguracao)
	configurarVariaveis(caminhoArquivoConfiguracao)

	e := rotas.GerarEcho()

	e.Server.WriteTimeout = 30 * time.Second
	if err := e.Start(fmt.Sprintf(":%d", PortaApi)); err != nil {
		logger.Logger().Fatal("Ocorreu um erro ao startar a api", err)
	}
}

// configurarVariaveis: configura as variaveis da api
func configurarVariaveis(caminhoArquivoConfiguracao string) {
	buscarPortaApi(caminhoArquivoConfiguracao)
}

// buscarPortaApi: busca a porta da API no arquivo de configuração
func buscarPortaApi(caminhoArquivoConfiguracao string) {
	porta, err := strconv.Atoi(os.Getenv("PortaApi"))
	if err != nil {
		PortaApi = PortaApiPadrao
		logger.Logger().Error("Ocorreu um erro ao converter a porta da api em string para int. Foi utilizada a porta padrão", err)
	} else {
		PortaApi = porta
	}
}
