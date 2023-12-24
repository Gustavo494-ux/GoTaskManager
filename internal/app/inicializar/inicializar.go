package inicializar

import (
	"GoTaskManager/internal/app/api"
	"GoTaskManager/pkg/routines/configuracoes"
)

// Inicializar: realiza todas as configurações para a inicialização do projeto
func Inicializar() {
	InicializarLogger()
	InicializarAPI()
}

// InicializarParaTestes: realiza configurações para execução dos testes
func InicializarParaTestes() {
	InicializarLogger()
}

// InicializarLogger: realize toda a configuração necessaria para utilização do logger
func InicializarLogger() {
	configuracoes.ConfigurarLogger("\\internal\\app\\config\\logger.yaml")
}

// InicializarAPI: realize toda a configuração necessaria para utilização da API
func InicializarAPI() {
	api.Configurar()
}
