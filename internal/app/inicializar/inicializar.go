package inicializar

import (
	"GoTaskManager/pkg/pacotes/manipuladorDeArquivos"
	"GoTaskManager/pkg/routines/configuracoes"
	"fmt"
	"os"
	"path/filepath"
)

var (
	NomeArquivoConfiguracaoPrincipal    = "arquivosConfiguracao.yaml"
	CaminhoArquivoConfiguracaoPrincipal string
)

func init() {
	DefinirDiretorioRaiz()
	CaminhoArquivoConfiguracaoPrincipal = filepath.Join(configuracoes.RetornarDiretorioRoot(), NomeArquivoConfiguracaoPrincipal)
}

// Inicializar: realiza todas as configurações para a inicialização do projeto
func Inicializar() {
	InicializarLogger()
	InicializarAPI()
}

// Carrega e define o diretorio raiz onde for necessario
func DefinirDiretorioRaiz() {
	diretorioRaiz := CarregarDiretorioRaiz()
	configuracoes.DefinirDiretorioRoot(diretorioRaiz)
	manipuladorDeArquivos.DefinirDiretorioRaiz(diretorioRaiz)
}

// InicializarParaTestes: realiza configurações para execução dos testes
func InicializarParaTestes() {
	InicializarLogger()
}

// InicializarLogger: realize toda a configuração necessaria para utilização do logger
func InicializarLogger() {
	configuracoes.ConfigurarLogger(configuracoes.BuscarParametroArquivoConfiguracao(CaminhoArquivoConfiguracaoPrincipal, "CaminhoArquivoLogger"))
}

// InicializarAPI: realize toda a configuração necessaria para utilização da API
func InicializarAPI() {
	configuracoes.ConfigurarApi(configuracoes.BuscarParametroArquivoConfiguracao(CaminhoArquivoConfiguracaoPrincipal, "CaminhoArquivoApi"))
}

// CarregarDiretorioRaiz: define o diretorio no qual o executavel está
func CarregarDiretorioRaiz() string {
	caminhoArquivo, _ := os.Getwd()
	caminhoArquivo, _ = manipuladorDeArquivos.ObterDiretorioDoArquivo(caminhoArquivo, NomeArquivoConfiguracaoPrincipal)
	fmt.Println(caminhoArquivo)

	return caminhoArquivo
}
