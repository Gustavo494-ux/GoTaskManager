package inicializar

import (
	"GoTaskManager/internal/app/inicializar/inicializarinternal"
	"GoTaskManager/pkg/routines/inicializarpkg"
)

func Inicializar() {
	inicializarpkg.Inicializar()
	inicializarinternal.Inicializar()
}

// InicializarParaTestes: realiza configurações para execução dos testes
func InicializarParaTestes() {
	inicializarpkg.InicializarParaTestes()
	inicializarinternal.InicializarParaTestes()
}
