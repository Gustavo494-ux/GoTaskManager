package configuracoes

import (
	"GoTaskManager/pkg/pacotes/GerenciadorArquivosConfig"
	"GoTaskManager/pkg/pacotes/authentication"
	"log"
)

// ConfigurarAutenticacao: configura a autenticacao
func ConfigurarAutenticacao(CaminhoRelativoArquivoConfiguracao string) {
	caminhoArquivoConfiguracao := FormatarCaminhoArquivoConfiguracao(CaminhoRelativoArquivoConfiguracao)
	PreencherVariaveisAutenteicacao(caminhoArquivoConfiguracao)
}

// PreencherVariaveisLog: Carrega os dados nas váriaveis
func PreencherVariaveisAutenteicacao(caminhoArquivoConfiguracao string) {
	authentication.DefinirSecretKey(buscarSecretKey(caminhoArquivoConfiguracao))
}

// buscarSecretKey: busca
func buscarSecretKey(caminhoArquivoConfiguracao string) string {
	secretKey, err := GerenciadorArquivosConfig.NovoArquivoConfig(caminhoArquivoConfiguracao).Ler().ObterValorParametro("CHAVE_SECRETA_JWT").String()
	if err != nil {
		log.Fatal("Ocorreu um erro ao buscar a SECRETA_JWT no arquivo de configuração de autenticação", err)
	}

	return secretKey
}
