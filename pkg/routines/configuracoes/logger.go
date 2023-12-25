package configuracoes

import (
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
	logger.FormatoDataHora = BuscarParametroArquivoConfiguracao(caminhoArquivoConfiguracao, "FormatoDataHora")
}

// ConfigurarCaminhoArquivoLog: cria o  arquivo de log caso não exista
func CriarArquivoLogSeNaoExistir(caminhoArquivoConfiguracao string) string {
	var caminhoArquivoLog, caminhoCompleto string
	caminhoArquivoLog = BuscarParametroArquivoConfiguracao(caminhoArquivoConfiguracao, "CaminhoArquivoLog")
	if len(caminhoArquivoLog) == 0 {
		caminhoCompleto = filepath.Join(diretorioRoot, CaminhoRelativoPadraoArquivoLog)
	} else {
		caminhoCompleto = caminhoArquivoLog
	}
	var err error

	err = manipuladorDeArquivos.CriarDiretorioOuArquivoSeNaoExistir(caminhoCompleto)
	if err != nil {
		log.Fatal("Ocorreu um erro ao criar o arquivo de log", err)
	}

	caminhoArquivoLog, err = manipuladorDeArquivos.ObterCaminhoAbsolutoOuConcatenadoComRaiz(caminhoCompleto)
	if err != nil {
		log.Fatal("Ocorreu um erro ao buscar o CaminhoArquivoLog", err)
	}

	caminhoArquivoLog = manipuladorDeArquivos.FormatarCaminho(caminhoArquivoLog)
	return caminhoArquivoLog
}
