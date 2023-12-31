package configuracoes

import (
	"GoTaskManager/pkg/pacotes/logger"
	"os"
)

const (
	CaminhoRelativoPadraoArquivoLog = "Logs/Logs.log"
)

// ConfigurarLogger: configura o Log
func ConfigurarLogger(CaminhoArquivoLog string) {
	caminhoFormatado := PrepararCaminhoArquivo(CaminhoArquivoLog)
	PreencherVariaveisLog(caminhoFormatado)
}

// PreencherVariaveisLog: Carrega os dados nas váriaveis
func PreencherVariaveisLog(CaminhoArquivoLog string) {
	logger.CaminhoArquivoLog = CaminhoArquivoLog
	logger.DiretorioRoot = diretorioRoot
	logger.FormatoDataHora = os.Getenv("FormatoDataHora")
}
