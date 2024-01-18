package configuracoes

import (
	"log"

	"github.com/Gustavo494-ux/PacotesGolang/manipuladorDeArquivos"
)

var (
	diretorioRoot string
)

func init() {

}

// PrepararCaminhoArquivo: retorna o caminho absoluto do arquivo de configurações
func PrepararCaminhoArquivo(caminho string) (caminhoArquivoConfiguracao string) {
	caminhoArquivoConfiguracao, err := manipuladorDeArquivos.ObterCaminhoAbsolutoOuConcatenadoComRaiz(caminho)
	if err != nil {
		log.Fatal("Ocorreu um erro ao preparar o caminho do arquivo "+caminho, err)
	}
	return
}

// DefinirDiretorioRoot: Define o caminho do diretorio root
func DefinirDiretorioRoot(diretorio string) {
	diretorioRoot = diretorio
}

// RetornarDiretorioRoot: retorna o caminho do diretorio root
func RetornarDiretorioRoot() string {
	return diretorioRoot
}
