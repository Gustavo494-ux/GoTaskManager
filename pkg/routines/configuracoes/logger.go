package configuracoes

import (
	"GoTaskManager/pkg/pacotes/logger"
	tipo "GoTaskManager/pkg/pacotes/tipos"
	"os"
)

const (
	CaminhoRelativoPadraoArquivoLog = "Logs/Logs.log"
)

// ConfigurarLogger: configura o Log
func ConfigurarLogger(CaminhoArquivoLog string) {
	if len(CaminhoArquivoLog) == 0 {
		CaminhoArquivoLog = CaminhoRelativoPadraoArquivoLog
	}

	caminhoFormatado := PrepararCaminhoArquivo(tipo.Coalesce().Str(CaminhoArquivoLog, CaminhoRelativoPadraoArquivoLog))
	PreencherVariaveisLog(caminhoFormatado)
}

// PreencherVariaveisLog: Carrega os dados nas váriaveis
func PreencherVariaveisLog(CaminhoArquivoLog string) {
	logger.CaminhoArquivoLog = CaminhoArquivoLog
	logger.DiretorioRoot = diretorioRoot
	logger.FormatoDataHora = os.Getenv("FormatoDataHora")
}
