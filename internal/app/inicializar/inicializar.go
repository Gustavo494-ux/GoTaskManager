package inicializar

import (
	"GoTaskManager/internal/app/api"
	"GoTaskManager/internal/app/migracao"
	"GoTaskManager/pkg/pacotes/logger"
	"GoTaskManager/pkg/pacotes/manipuladorDeArquivos"
	"GoTaskManager/pkg/routines/configuracoes"
	"os"
	"path/filepath"
	"strconv"
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
	InicializarBancoDeDadosPrincipal()
	InicializarTabelasBancoDeDados()
	InicializarAPI()
}

// InicializarParaTestes: realiza configurações para execução dos testes
func InicializarParaTestes() {
	InicializarLogger()
	InicializarBancoDeDadosTeste()
	InicializarAPI()
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

// InicializarAPI: realize toda a configuração necessaria para utilização da API
func InicializarAPI() {
	api.ConfigurarApi(os.Getenv("PortaApi"))
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
	porta, err := strconv.Atoi(os.Getenv("PORTA_DATABASE"))
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao converter a PORTA_DATABASE para string", err)
	}

	configuracoes.BancoProducao = configuracoes.ConfigurarNovoBanco(
		os.Getenv("HOST_DATABASE"),
		os.Getenv("NOME_DATABASE"),
		os.Getenv("USUARIO_DATABASE"),
		os.Getenv("SENHA_DATABASE"),
		os.Getenv("NOME_DRIVER_DATABASE"),
		os.Getenv("SSLMODE_DATABASE"),
		porta,
	)
}

// InicializarBancoDeDadosTeste: inicializa o banco de dados para uso dos testes
func InicializarBancoDeDadosTeste() {
	porta, err := strconv.Atoi(os.Getenv("PORTA_DATABASE_TESTE"))
	if err != nil {
		logger.Logger().Fatal("Ocorreu um erro ao converter a PORTA_DATABASE_TESTE para string", err)
	}

	configuracoes.BancoProducao = configuracoes.ConfigurarNovoBanco(
		os.Getenv("HOST_DATABASE_TESTE "),
		os.Getenv("NOME_DATABASE_TESTE "),
		os.Getenv("USUARIO_DATABASE_TESTE "),
		os.Getenv("SENHA_DATABASE_TESTE "),
		os.Getenv("NOME_DRIVER_DATABASE_TESTE "),
		os.Getenv("SSLMODE_DATABASE_TESTE "),
		porta,
	)
}

// InicializarTabelasBancoDeDados: cria as tabelas do banco de dados automaticamente
func InicializarTabelasBancoDeDados() {
	migracao.CriacaoAutomaticaTabelas()
}
