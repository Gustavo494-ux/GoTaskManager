package configuracoes

import (
	"GoTaskManager/pkg/pacotes/GerenciadorArquivosConfig"
	"GoTaskManager/pkg/pacotes/logger"
	"GoTaskManager/pkg/pacotes/manipuladorDeArquivos"
	"log"
	"path/filepath"
)

const (
	CaminhoRelativoPadraoArquivoLog = "\\Logs\\Logs.log"
)

// ConfigurarLogger: configura o Log
func ConfigurarLogger(CaminhoRelativoArquivoConfiguracao string) {
	caminhoArquivoConfiguracao := FormatarCaminhoArquivoConfiguracao(CaminhoRelativoArquivoConfiguracao)
	PreencherVariaveisLog(caminhoArquivoConfiguracao)
}

// PreencherVariaveisLog: Carrega os dados nas váriaveis
func PreencherVariaveisLog(caminhoArquivoConfiguracao string) {
	logger.CaminhoArquivoLog = CriarArquivoLogSeNaoExistir(caminhoArquivoConfiguracao)
	logger.DiretorioRoot = diretorioRoot
	logger.FormatoDataHora = buscarFormatoDataHora(caminhoArquivoConfiguracao)
}

// ConfigurarCaminhoArquivoLog: cria o  arquivo de log caso não exista
func CriarArquivoLogSeNaoExistir(caminhoArquivoConfiguracao string) string {
	var caminhoArquivoLog string
	caminhoArquivoLog = buscarCaminhoArquivoLog(caminhoArquivoConfiguracao)
	if len(caminhoArquivoLog) == 0 {
		caminhoCompleto := filepath.Join(diretorioRoot, CaminhoRelativoPadraoArquivoLog)

		var err error

		manipuladorDeArquivos.CriarDiretorioOuArquivoSeNaoExistir(caminhoCompleto)
		if err != nil {
			log.Fatal("Ocorreu um erro ao criar o arquivo de log", err)
		}

		manipuladorDeArquivos.CriarDiretorioOuArquivoSeNaoExistir(caminhoCompleto)
		caminhoArquivoLog, err = manipuladorDeArquivos.ObterCaminhoAbsolutoOuConcatenadoComRaiz(caminhoCompleto)
		if err != nil {
			log.Fatal("Ocorreu um erro ao buscar o CaminhoArquivoLog", err)
		}
	}
	caminhoArquivoLog = manipuladorDeArquivos.FormatarCaminho(caminhoArquivoLog)
	return caminhoArquivoLog
}

// buscarCaminhoArquivoLog: retorna o caminho do arquivo de log formatado
func buscarCaminhoArquivoLog(caminhoArquivoConfiguracao_Logger string) string {
	caminhoArquivoLog, err := GerenciadorArquivosConfig.NovoArquivoConfig(
		manipuladorDeArquivos.FormatarCaminho(caminhoArquivoConfiguracao_Logger)).
		Ler().
		ObterValorParametro("CaminhoArquivoLog").
		String()
	if err != nil {
		log.Fatal("Ocorreu um erro ao buscar o caminho do arquivo de logger", err)
	}

	return caminhoArquivoLog
}

// buscarFormatoDataHora
func buscarFormatoDataHora(caminhoArquivoConfiguracao_Logger string) string {
	formatoDataHora, err := GerenciadorArquivosConfig.NovoArquivoConfig(caminhoArquivoConfiguracao_Logger).Ler().ObterValorParametro("FormatoDataHora").String()
	if err != nil {
		log.Fatal("Ocorreu um erro ao formato de data e hora do logger", err)
	}

	return formatoDataHora
}
