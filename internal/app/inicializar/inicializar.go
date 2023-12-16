package inicializar

import "GoTaskManager/pkg/routines/configurations"

// Inicializar: realiza todas as configurações para a inicialização do projeto
func Inicializar() {
	inicializarLogger()
}

// InicializarParaTestes: realiza configurações para execução dos testes
func InicializarParaTestes() {
	inicializarLogger()
}

// inicializarLogger: realize toda a configuração necessaria para utilização do logger
func inicializarLogger() {
	configurations.ConfigurarLogger("\\internal\\app\\config\\logger.yaml")
}
