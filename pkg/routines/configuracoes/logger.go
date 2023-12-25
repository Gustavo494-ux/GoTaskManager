package configuracoes

import (
	"GoTaskManager/pkg/pacotes/logger"
	"GoTaskManager/pkg/pacotes/manipuladorDeArquivos"
	"log"
	"os"
	"path/filepath"
)

const (
	CaminhoRelativoPadraoArquivoLog = "\\Logs\\Logs.log"
)

// ConfigurarLogger: configura o Log
func ConfigurarLogger(CaminhoArquivoLog string) {
	caminhoFormatado := FormatarCaminhoArquivoConfiguracao(CaminhoArquivoLog)
	PreencherVariaveisLog(caminhoFormatado)
}

// PreencherVariaveisLog: Carrega os dados nas váriaveis
func PreencherVariaveisLog(CaminhoArquivoLog string) {
	logger.CaminhoArquivoLog = CriarArquivoLogSeNaoExistir(CaminhoArquivoLog)
	logger.DiretorioRoot = diretorioRoot
	logger.FormatoDataHora = os.Getenv("FormatoDataHora")
}

// ConfigurarCaminhoArquivoLog: cria o  arquivo de log caso não exista
func CriarArquivoLogSeNaoExistir(caminhoArquivoLog string) string {
	var caminhoCompleto string
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
