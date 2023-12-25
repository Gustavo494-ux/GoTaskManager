package env

import (
	"GoTaskManager/pkg/pacotes/logger"

	"github.com/joho/godotenv"
)

var (
	caminhoArquivoEnv string
)

// CarregarDotEnv: carrega o arquivo .env
func CarregarDotEnv() {
	err := godotenv.Load(caminhoArquivoEnv)
	if err != nil {
		logger.Logger().Fatal("Ocorreu um erro ao carregar o arquivo env", err, caminhoArquivoEnv)
	}
}

// DefinirCaminhoArquivoEnv: Define o caminho do arquivo env
func DefinirCaminhoArquivoEnv(caminho string) {
	caminhoArquivoEnv = caminho
}

// RetornarDiretorioRoot: retorna o caminho do arquivo env
func RetornarCaminhoArquivoEnv() string {
	return caminhoArquivoEnv
}
