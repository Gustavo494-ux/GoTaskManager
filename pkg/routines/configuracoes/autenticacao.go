package configuracoes

import (
	"GoTaskManager/pkg/pacotes/authentication"
	"os"
)

// ConfigurarAutenticacao: configura a autenticacao
func ConfigurarAutenticacao(CaminhoRelativoArquivoConfiguracao string) {
	caminhoArquivoConfiguracao := PrepararCaminhoArquivo(CaminhoRelativoArquivoConfiguracao)
	PreencherVariaveisAutenteicacao(caminhoArquivoConfiguracao)
}

// PreencherVariaveisLog: Carrega os dados nas váriaveis
func PreencherVariaveisAutenteicacao(caminhoArquivoConfiguracao string) {
	authentication.DefinirSecretKey(buscarSecretKey(caminhoArquivoConfiguracao))
}

// buscarSecretKey: busca
func buscarSecretKey(caminhoArquivoConfiguracao string) string {
	return os.Getenv("CHAVE_SECRETA_JWT")
}
