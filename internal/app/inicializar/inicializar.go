package inicializar

import (
	"GoTaskManager/pkg/pacotes/manipuladorDeArquivos"
	"GoTaskManager/pkg/routines/configuracoes"
	"os"
	"path/filepath"
)

var (
	nomeArquivoConfiguracao = "app.env"
	diretorioRaiz           string
)

func init() {
	DefinirDiretorioRaiz()
	InicializarDotEnv()
}

// Inicializar: realiza todas as configurações para a inicialização do projeto
func Inicializar() {
	InicializarLogger()
	InicializarAPI()
}

// Carrega e define o diretorio raiz onde for necessario
func DefinirDiretorioRaiz() {
	diretorioRaiz = CarregarDiretorioRaiz()
	configuracoes.DefinirDiretorioRoot(diretorioRaiz)
	manipuladorDeArquivos.DefinirDiretorioRaiz(diretorioRaiz)
}

// InicializarParaTestes: realiza configurações para execução dos testes
func InicializarParaTestes() {
	InicializarLogger()
}

// InicializarLogger: realize toda a configuração necessaria para utilização do logger
func InicializarLogger() {
	configuracoes.ConfigurarLogger(os.Getenv("CaminhoArquivoLogger"))
}

// InicializarAPI: realize toda a configuração necessaria para utilização da API
func InicializarAPI() {
	configuracoes.ConfigurarApi(os.Getenv("CaminhoArquivoApi"))
}

// InicializarAPI: realize toda a configuração necessaria para utilização do env
func InicializarDotEnv() {
	configuracoes.ConfigurarEnv(filepath.Join(diretorioRaiz, "/", nomeArquivoConfiguracao))
}

// CarregarDiretorioRaiz: define o diretorio no qual o executavel está
func CarregarDiretorioRaiz() string {
	caminhoArquivo, _ := os.Getwd()
	caminhoArquivo, _ = manipuladorDeArquivos.ObterDiretorioDoArquivo(caminhoArquivo, nomeArquivoConfiguracao)

	return caminhoArquivo
}
