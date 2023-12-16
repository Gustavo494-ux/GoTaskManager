package configurations

import (
	"GoTaskManager/pkg/pacotes/GerenciadorArquivosConfig"
	"GoTaskManager/pkg/pacotes/logger"
	"GoTaskManager/pkg/pacotes/manipuladorDeArquivos"
	"log"
	"path/filepath"
)

var (
	// Parametro
	CaminhoRelativoArquivoConfiguracao string

	// Preenchidas automaticamente
	DiretorioRoot              string
	CaminhoArquivoConfiguracao string

	// Preenchidas por arquivo de configuração
	CaminhoArquivoLog string
	FormatoDataHora   string
)

const (
	CaminhoRelativoPadraoArquivoLog = "\\Logs\\Logs.log"
)

func ConfigurarLogger(caminhoRelativoArquivoConfiguracao string) {
	CaminhoRelativoArquivoConfiguracao = caminhoRelativoArquivoConfiguracao
	buscarDiretorioRoot()
	ConfigurarCaminhoArquivoConfiguracao()
	ConfigurarCaminhoArquivoLog()
	buscarFormatoDataHora()
	PreencherVariaveisLog()
}

func PreencherVariaveisLog() {
	logger.CaminhoArquivoLog = CaminhoArquivoLog
	logger.DiretorioRoot = DiretorioRoot
	logger.FormatoDataHora = FormatoDataHora
}

func ConfigurarCaminhoArquivoConfiguracao() {
	var err error
	CaminhoArquivoConfiguracao, err = manipuladorDeArquivos.ObterCaminhoAbsolutoOuConcatenadoComRaiz(filepath.Join(DiretorioRoot, CaminhoRelativoArquivoConfiguracao))
	if err != nil {
		if err != nil {
			log.Fatal("Ocorreu um erro ao buscar o CaminhoArquivoConfiguracao", err)
		}
	}
}

func ConfigurarCaminhoArquivoLog() {
	buscarCaminhoArquivoLog()
	if len(CaminhoArquivoLog) == 0 {
		caminhoCompleto := filepath.Join(DiretorioRoot, CaminhoRelativoPadraoArquivoLog)

		var err error

		manipuladorDeArquivos.CriarDiretorioOuArquivoSeNaoExistir(caminhoCompleto)
		if err != nil {
			log.Fatal("Ocorreu um erro ao criar o arquivo de log", err)
		}

		manipuladorDeArquivos.CriarDiretorioOuArquivoSeNaoExistir(caminhoCompleto)
		CaminhoArquivoLog, err = manipuladorDeArquivos.ObterCaminhoAbsolutoOuConcatenadoComRaiz(caminhoCompleto)
		if err != nil {
			log.Fatal("Ocorreu um erro ao buscar o CaminhoArquivoLog", err)
		}

		CaminhoArquivoLog = manipuladorDeArquivos.FormatarCaminho(CaminhoArquivoLog)

	}
}

func buscarDiretorioRoot() {
	var err error
	DiretorioRoot, err = manipuladorDeArquivos.BuscarDiretorioRootRepositorio()
	if err != nil {
		log.Fatal("Diretorio root do repositorio não encontrado erro: ", err)
	}
}

func buscarCaminhoArquivoLog() {
	var err error
	CaminhoArquivoLog, err = GerenciadorArquivosConfig.NovoArquivoConfig(
		manipuladorDeArquivos.FormatarCaminho(CaminhoArquivoConfiguracao)).
		Ler().
		ObterValorParametro("CaminhoArquivoLog").
		String()
	if err != nil {
		log.Fatal("Ocorreu um erro ao buscar o caminho do arquivo de logger", err)
	}
}

func buscarFormatoDataHora() {
	var err error
	FormatoDataHora, err = GerenciadorArquivosConfig.NovoArquivoConfig(CaminhoArquivoConfiguracao).Ler().ObterValorParametro("FormatoDataHora").String()
	if err != nil {
		log.Fatal("Ocorreu um erro ao formato de data e hora do logger", err)
	}
}
