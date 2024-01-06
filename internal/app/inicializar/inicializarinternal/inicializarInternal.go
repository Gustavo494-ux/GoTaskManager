package inicializarinternal

import (
	"GoTaskManager/internal/app/api"
	"GoTaskManager/internal/app/migracao"
	"os"
)

// Inicializar: realiza todas as configurações nos pacotes INTERNAL para a inicialização do projeto
func Inicializar() {
	InicializarTabelasBancoDeDados()
	InicializarAPI()
}

// InicializarParaTestes: realiza configurações para execução dos testes
func InicializarParaTestes() {
	InicializarTabelasBancoDeDados()
	InicializarAPI()
}

// InicializarAPI: realize toda a configuração necessaria para utilização da API
func InicializarAPI() {
	api.ConfigurarApi(os.Getenv("PortaApi"))
}

// InicializarTabelasBancoDeDados: cria as tabelas do banco de dados automaticamente
func InicializarTabelasBancoDeDados() {
	migracao.CriacaoAutomaticaTabelas()
}
