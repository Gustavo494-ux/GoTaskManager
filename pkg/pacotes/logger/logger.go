package logger

import (
	"GoTaskManager/pkg/pacotes/GerenciadordeJson"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

type NivelLog zerolog.Level

type LoggerType struct {
	log             zerolog.Logger
	mensagem        string
	SalvarEmArquivo bool
}

var (
	CaminhoArquivoLog string
	DiretorioRoot     string
	FormatoDataHora   string
	DiretorioRaiz     string
)

const (
	NivelLog_Debug        NivelLog = NivelLog(zerolog.DebugLevel)
	NivelLog_Desabilitado NivelLog = NivelLog(zerolog.Disabled)
	NivelLog_Informacoes  NivelLog = NivelLog(zerolog.InfoLevel)
	NivelLog_Erro         NivelLog = NivelLog(zerolog.ErrorLevel)

	NivelLog_Panico       NivelLog = NivelLog(zerolog.PanicLevel)
	NivelLog_Rastreamento NivelLog = NivelLog(zerolog.TraceLevel)
)

// Logger cria uma instância de logger
func Logger() *LoggerType {
	var logger LoggerType
	logger.init()
	return &logger
}

// init: realiza a configuração necessária para o pacote funcionar
func (logger *LoggerType) init() {
	logger.configurarLog(NivelLog_Rastreamento)
}

// configurarLog: realiza a configuração básica para o log funcionar
func (logger *LoggerType) configurarLog(nivelLog NivelLog) {
	zerolog.TimeFieldFormat = FormatoDataHora

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if nivelLog < -1 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	caminho, err := ObterCaminhoAbsolutoOuConcatenadoComRaiz(CaminhoArquivoLog)
	if err != nil {
		log.Error("Ocorreu um erro ao montar o caminho do diretorio raiz", err, caminho)
	}

	arquivoLog, err := CarregarArquivo(caminho)
	if err != nil {
		log.Error("Erro ao carregar arquivo de log", err)
	} else {
		logger.SalvarEmArquivo = true
	}

	logger.log = zerolog.New(arquivoLog).With().Timestamp().Logger()
}

// Fatal: cria um log de erro fatal
func (logger *LoggerType) Fatal(mensagem string, err error, dados ...interface{}) {
	logger.mensagem = mensagem
	fmt.Println(mensagem)

	if logger.SalvarEmArquivo {
		logger.log.
			Fatal().
			Caller(1).
			Err(err).
			Str("Dados Adicionais", logger.converterSliceDadosParaJsonString(dados)).
			Msg(mensagem)
	}
}

// Error: cria um log de erro
func (logger *LoggerType) Error(mensagem string, err error, dados ...interface{}) {
	logger.mensagem = mensagem
	fmt.Println(mensagem)
	if logger.SalvarEmArquivo {
		logger.log.
			Error().
			Caller(1).
			Err(err).
			Str("Dados Adicionais", logger.converterSliceDadosParaJsonString(dados)).
			Msg(mensagem)
	}
}

// Alerta: cria um log de Alerta
func (logger *LoggerType) Alerta(mensagem string, dados ...interface{}) {
	if logger.SalvarEmArquivo {
		logger.mensagem = mensagem
		fmt.Println(mensagem)
		logger.log.
			Warn().
			Caller(1).
			Str("Dados Adicionais", logger.converterSliceDadosParaJsonString(dados)).
			Msg(mensagem)
	}
}

// Info: cria um log de informação
func (logger *LoggerType) Info(mensagem string, dados ...interface{}) {
	logger.mensagem = mensagem
	fmt.Println(mensagem)
	if logger.SalvarEmArquivo {
		logger.log.
			Info().
			Caller(1).
			Str("Dados Adicionais", logger.converterSliceDadosParaJsonString(dados)).
			Msg(mensagem)
	}
}

// Debug: cria um log de Debug
func (logger *LoggerType) Debug(mensagem string, dados ...interface{}) {
	logger.mensagem = mensagem
	fmt.Println(mensagem)
	if logger.SalvarEmArquivo {
		logger.log.
			Debug().
			Caller(1).
			Str("Dados Adicionais", logger.converterSliceDadosParaJsonString(dados)).
			Msg(mensagem)
	}
}

// Rastreamento: cria um log de rastreamento
func (logger *LoggerType) Rastreamento(mensagem string, dados ...interface{}) {
	logger.mensagem = mensagem
	fmt.Println(mensagem)
	if logger.SalvarEmArquivo {
		logger.log.
			Trace().
			Caller(1).
			Str("Dados Adicionais", logger.converterSliceDadosParaJsonString(dados)).
			Msg(mensagem)
	}
}

// converterSliceDadosParaJsonString: converte uma interface para jsonString
func (logger *LoggerType) converterSliceDadosParaJsonString(dados ...interface{}) (jsonString string) {
	var dado string
	var err error
	for _, Valor := range dados {
		dado, err = GerenciadordeJson.InterfaceParaJsonString(Valor)
		if err != nil {
			log.Error(err)
			return
		}
		jsonString += dado
	}
	return
}

// converterSliceDadosParaJsonString: converte uma interface para jsonString
func (logger *LoggerType) RetornarMensagem() string {
	return logger.mensagem
}

// ObterCaminhoAbsolutoOuConcatenadoComRaiz retorna o caminho absoluto ou o ultimo diretorio raiz + ultimo diretorio do parametro caminho + nome do arquivo
func ObterCaminhoAbsolutoOuConcatenadoComRaiz(caminho string) (string, error) {
	caminho = filepath.FromSlash(caminho)
	if filepath.IsAbs(caminho) {
		return caminho, nil
	}

	// Obter o último diretório e nome do arquivo do caminho
	dir, nomeArquivo := filepath.Split(caminho)
	dir = strings.TrimSuffix(dir, string(filepath.Separator))
	nomeArquivo = strings.TrimPrefix(nomeArquivo, string(filepath.Separator))

	// Concatenar o último diretório e nome do arquivo com o caminho raiz
	caminhoAbsoluto := filepath.Join(DiretorioRaiz, dir, nomeArquivo)
	return caminhoAbsoluto, nil
}

// CarregarArquivo abre um arquivo existente para leitura e adição de logs
func CarregarArquivo(caminho string) (file *os.File, err error) {
	if err = os.MkdirAll(strings.ReplaceAll(filepath.Dir(caminho), "\\", "/"), os.ModePerm); err != nil {
		return
	}

	arquivo, err := os.OpenFile(caminho, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return arquivo, nil
}
