package inicializar

import (
	"GoTaskManager/pkg/routines/configuracoes"
)

// Inicializar: realiza todas as configurações para a inicialização do projeto
func Inicializar() {
	InicializarLogger()
}

// InicializarParaTestes: realiza configurações para execução dos testes
func InicializarParaTestes() {
	InicializarLogger()
}

// InicializarLogger: realize toda a configuração necessaria para utilização do logger
func InicializarLogger() {
	configuracoes.ConfigurarLogger("\\internal\\app\\config\\logger.yaml")
}
