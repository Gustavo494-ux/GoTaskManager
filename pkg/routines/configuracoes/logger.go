package configuracoes

import (
	"GoTaskManager/pkg/pacotes/GerenciadorArquivosConfig"
	"GoTaskManager/pkg/pacotes/logger"
	"GoTaskManager/pkg/pacotes/manipuladorDeArquivos"
	"log"
	"path/filepath"
)

var (
	// Parametro
	caminhoRelativoArquivoConfiguracao_Logger string

	// Preenchidas automaticamente
	caminhoArquivoConfiguracao_Logger string

	// Preenchidas por arquivo de configuração
	caminhoArquivoLog_Logger string
	formatoDataHora_Logger   string
)

const (
	CaminhoRelativoPadraoArquivoLog = "\\Logs\\Logs.log"
)

func ConfigurarLogger(CaminhoRelativoArquivoConfiguracao string) {
	caminhoRelativoArquivoConfiguracao_Logger = CaminhoRelativoArquivoConfiguracao
	buscarDiretorioRoot()
	ConfigurarCaminhoArquivoConfiguracao()
	ConfigurarCaminhoArquivoLog()
	buscarFormatoDataHora()
	PreencherVariaveisLog()
}

func PreencherVariaveisLog() {
	logger.CaminhoArquivoLog = caminhoArquivoLog_Logger
	logger.DiretorioRoot = diretorioRoot
	logger.FormatoDataHora = formatoDataHora_Logger
}

func ConfigurarCaminhoArquivoConfiguracao() {
	var err error
	caminhoArquivoConfiguracao_Logger, err = manipuladorDeArquivos.ObterCaminhoAbsolutoOuConcatenadoComRaiz(
		filepath.Join(diretorioRoot, caminhoRelativoArquivoConfiguracao_Logger))
	if err != nil {
		if err != nil {
			log.Fatal("Ocorreu um erro ao buscar o CaminhoArquivoConfiguracao", err)
		}
	}
}

func ConfigurarCaminhoArquivoLog() {
	buscarCaminhoArquivoLog()
	if len(caminhoArquivoLog_Logger) == 0 {
		caminhoCompleto := filepath.Join(diretorioRoot, CaminhoRelativoPadraoArquivoLog)

		var err error

		manipuladorDeArquivos.CriarDiretorioOuArquivoSeNaoExistir(caminhoCompleto)
		if err != nil {
			log.Fatal("Ocorreu um erro ao criar o arquivo de log", err)
		}

		manipuladorDeArquivos.CriarDiretorioOuArquivoSeNaoExistir(caminhoCompleto)
		caminhoArquivoLog_Logger, err = manipuladorDeArquivos.ObterCaminhoAbsolutoOuConcatenadoComRaiz(caminhoCompleto)
		if err != nil {
			log.Fatal("Ocorreu um erro ao buscar o CaminhoArquivoLog", err)
		}

		caminhoArquivoLog_Logger = manipuladorDeArquivos.FormatarCaminho(caminhoArquivoLog_Logger)

	}
}

func buscarDiretorioRoot() {
	var err error
	diretorioRoot, err = manipuladorDeArquivos.BuscarDiretorioRootRepositorio()
	if err != nil {
		log.Fatal("Diretorio root do repositorio não encontrado erro: ", err)
	}
}

func buscarCaminhoArquivoLog() {
	var err error
	caminhoArquivoLog_Logger, err = GerenciadorArquivosConfig.NovoArquivoConfig(
		manipuladorDeArquivos.FormatarCaminho(caminhoArquivoConfiguracao_Logger)).
		Ler().
		ObterValorParametro("CaminhoArquivoLog").
		String()
	if err != nil {
		log.Fatal("Ocorreu um erro ao buscar o caminho do arquivo de logger", err)
	}
}

func buscarFormatoDataHora() {
	var err error
	formatoDataHora_Logger, err = GerenciadorArquivosConfig.NovoArquivoConfig(caminhoArquivoConfiguracao_Logger).Ler().ObterValorParametro("FormatoDataHora").String()
	if err != nil {
		log.Fatal("Ocorreu um erro ao formato de data e hora do logger", err)
	}
}
