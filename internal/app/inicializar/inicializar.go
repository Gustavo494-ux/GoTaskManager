package inicializar

import (
	inicializarinternal "GoTaskManager/internal/app/inicializar/inicializarInternal"
	inicializarpkg "GoTaskManager/pkg/routines/inicializarPKG"
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
