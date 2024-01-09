package utils

import (
	"GoTaskManager/pkg/pacotes/encryption"
	"GoTaskManager/pkg/pacotes/logger"
)

// GerarHash: gera um hash a partir do parametro passado
func GerarHash(texto string) (s string) {
	s, err := encryption.GerarSHA512(texto)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao gerar o hash", err, texto)
	}
	return
}

// CompararHASH: verifica se o texto e a hash são equivalentes
func CompararHASH(hash, dadoTexto string) bool {
	return encryption.CompararSHA512(hash, dadoTexto) == nil
}
