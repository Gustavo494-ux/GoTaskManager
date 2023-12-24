package configuracoes

import (
	"GoTaskManager/internal/app/api/rotas"
	"GoTaskManager/pkg/pacotes/GerenciadorArquivosConfig"
	"GoTaskManager/pkg/pacotes/logger"
	"fmt"
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
func Configurar(CaminhoRelativoArquivoConfiguracao string) {
	caminhoArquivoConfiguracao := FormatarCaminhoArquivoConfiguracao(CaminhoRelativoArquivoConfiguracao)
	configurarVariaveis(caminhoArquivoConfiguracao)

	e := rotas.GerarEcho()

	e.Server.WriteTimeout = 30 * time.Second
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PortaApi)))
}

// configurarVariaveis: configura as variaveis da api
func configurarVariaveis(caminhoArquivoConfiguracao string) {
	buscarPortaApi(caminhoArquivoConfiguracao)
}

// buscarPortaApi: busca a porta da API no arquivo de configuração
func buscarPortaApi(caminhoArquivoConfiguracao string) {
	parametro, err := GerenciadorArquivosConfig.NovoArquivoConfig(caminhoArquivoConfiguracao).Ler().ObterValorParametro("PortaApi").String()
	if err != nil {
		logger.Logger().Error("Ocorreu um erro buscar a porta da api no arquivo de configuração da api", err)
		PortaApi = PortaApiPadrao
	}

	portaApi, err := strconv.Atoi(parametro)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro converter a porta da api recuperada do arquivo de configuração da api em int", err)
		PortaApi = PortaApiPadrao
	}

	if portaApi > 0 {
		PortaApi = portaApi
	} else {
		PortaApi = PortaApiPadrao
	}

}
