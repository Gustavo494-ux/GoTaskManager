package configuracoes

import (
	"GoTaskManager/pkg/pacotes/manipuladorDeArquivos"
	"log"
	"path/filepath"
)

var (
	diretorioRoot string
)

func init() {
	carregarDiretorioRoot()
}

// FormatarCaminhoArquivoConfiguracao: retorna o caminho absoluto do arquivo de configurações
func FormatarCaminhoArquivoConfiguracao(caminhoRelativoArquivoConfiguracao string) string {
	caminhoArquivoConfiguracao, err := manipuladorDeArquivos.ObterCaminhoAbsolutoOuConcatenadoComRaiz(
		filepath.Join(diretorioRoot, caminhoRelativoArquivoConfiguracao),
	)

	if err != nil {
		log.Fatal("Ocorreu um erro ao buscar o CaminhoArquivoConfiguracao", err)
	}

	return caminhoArquivoConfiguracao
}

// carregarDiretorioRoot: carrega na variavel diretorioRoot o caminho do diretorio raiz do repositorio
func carregarDiretorioRoot() {
	var err error
	diretorioRoot, err = manipuladorDeArquivos.BuscarDiretorioRootRepositorio()

	if err != nil {
		log.Fatal("Diretorio root do repositorio não encontrado erro: ", err)
	}
}
