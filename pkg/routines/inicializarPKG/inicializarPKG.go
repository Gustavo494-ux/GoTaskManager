package inicializarpkg

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

// Inicializar: realiza todas as configurações nos pacotes PKG para a inicialização do projeto
func Inicializar() {
	InicializarLogger()
	InicializarBancoDeDadosPrincipal()
}

// InicializarParaTestes: realiza configurações para execução dos testes
func InicializarParaTestes() {
	InicializarLogger()
	InicializarBancoDeDadosTeste()
}

// Carrega e define o diretorio raiz onde for necessario
func DefinirDiretorioRaiz() {
	diretorioRaiz = CarregarDiretorioRaiz()
	configuracoes.DefinirDiretorioRoot(diretorioRaiz)
	manipuladorDeArquivos.DefinirDiretorioRaiz(diretorioRaiz)
}

// InicializarLogger: realize toda a configuração necessaria para utilização do logger
func InicializarLogger() {
	configuracoes.ConfigurarLogger(os.Getenv("CaminhoArquivoLogger"))
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

// InicializarBancoDeDadosPrincipal: inicializa o banco de dados para uso
func InicializarBancoDeDadosPrincipal() {
	configuracoes.BancoProducao = configuracoes.ConfigurarNovoBanco(
		os.Getenv("HOST_DATABASE"),
		os.Getenv("NOME_DATABASE"),
		os.Getenv("USUARIO_DATABASE"),
		os.Getenv("SENHA_DATABASE"),
		os.Getenv("NOME_DRIVER_DATABASE"),
		os.Getenv("SSLMODE_DATABASE"),
		os.Getenv("PORTA_DATABASE"),
	)
}

// InicializarBancoDeDadosTeste: inicializa o banco de dados para uso dos testes
func InicializarBancoDeDadosTeste() {
	configuracoes.BancoProducao = configuracoes.ConfigurarNovoBanco(
		os.Getenv("HOST_DATABASE_TESTE "),
		os.Getenv("NOME_DATABASE_TESTE "),
		os.Getenv("USUARIO_DATABASE_TESTE "),
		os.Getenv("SENHA_DATABASE_TESTE "),
		os.Getenv("NOME_DRIVER_DATABASE_TESTE "),
		os.Getenv("SSLMODE_DATABASE_TESTE "),
		os.Getenv("PORTA_DATABASE_TESTE"),
	)
}
