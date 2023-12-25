package configuracoes

import (
	"GoTaskManager/pkg/pacotes/GerenciadorArquivosConfig"
	"GoTaskManager/pkg/pacotes/manipuladorDeArquivos"
	"log"
)

var (
	diretorioRoot string
)

func init() {

}

// FormatarCaminhoArquivoConfiguracao: retorna o caminho absoluto do arquivo de configurações
func FormatarCaminhoArquivoConfiguracao(caminhoRelativoArquivoConfiguracao string) string {
	caminhoArquivoConfiguracao, err := manipuladorDeArquivos.ObterCaminhoAbsolutoOuConcatenadoComRaiz(caminhoRelativoArquivoConfiguracao)

	if err != nil {
		log.Fatal("Ocorreu um erro ao buscar o CaminhoArquivoConfiguracao", err)
	}

	return manipuladorDeArquivos.FormatarCaminho(caminhoArquivoConfiguracao)
}

// DefinirDiretorioRoot: Define o caminho do diretorio root
func DefinirDiretorioRoot(diretorio string) {
	diretorioRoot = diretorio
}

// RetornarDiretorioRoot: retorna o caminho do diretorio root
func RetornarDiretorioRoot() string {
	return diretorioRoot
}

// BuscarParametroArquivoConfiguracao: busca e retorna o parametro de configuração do arquivo yaml
func BuscarParametroArquivoConfiguracao(caminhoArquivoConfiguracao string, chave string) string {
	caminhoArquivoConfiguracao = FormatarCaminhoArquivoConfiguracao(caminhoArquivoConfiguracao)

	parametroConfiguracao, err := GerenciadorArquivosConfig.NovoArquivoConfig(caminhoArquivoConfiguracao).
		Ler().
		ObterValorParametro(chave).
		String()

	if err != nil {
		log.Fatal("Ocorreu um erro ao buscar o parametro de configuração "+chave, err)
	}

	return parametroConfiguracao
}
