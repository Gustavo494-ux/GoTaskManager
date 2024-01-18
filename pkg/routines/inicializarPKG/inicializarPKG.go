package inicializarpkg

import (
	"GoTaskManager/pkg/routines/configuracoes"
	"os"
	"path/filepath"

	"github.com/Gustavo494-ux/PacotesGolang/manipuladorDeArquivos"
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
	configuracoes.ConfiguracaoBancoProducao = configuracoes.ConfigurarNovoBanco(
		os.Getenv("HOST_DATABASE"),
		os.Getenv("NOME_DATABASE"),
		os.Getenv("USUARIO_DATABASE"),
		os.Getenv("SENHA_DATABASE"),
		os.Getenv("NOME_DRIVER_DATABASE"),
		os.Getenv("SSLMODE_DATABASE"),
		os.Getenv("PORTA_DATABASE"),
	)

	configuracoes.BancoPrincipalGORM = configuracoes.ConfiguracaoBancoProducao.ConectarGorm()
}

// InicializarBancoDeDadosTeste: inicializa o banco de dados para uso dos testes
func InicializarBancoDeDadosTeste() {
	configuracoes.ConfiguracaoBancoTeste = configuracoes.ConfigurarNovoBanco(
		os.Getenv("HOST_DATABASE_TESTE"),
		os.Getenv("NOME_DATABASE_TESTE"),
		os.Getenv("USUARIO_DATABASE_TESTE"),
		os.Getenv("SENHA_DATABASE_TESTE"),
		os.Getenv("NOME_DRIVER_DATABASE_TESTE"),
		os.Getenv("SSLMODE_DATABASE_TESTE"),
		os.Getenv("PORTA_DATABASE_TESTE"),
	)
	configuracoes.BancoPrincipalGORM = configuracoes.ConfiguracaoBancoTeste.ConectarGorm()
}
