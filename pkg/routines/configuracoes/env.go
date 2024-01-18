package configuracoes

import "github.com/Gustavo494-ux/PacotesGolang/env"

//ConfigurarEnv: realiza a configuração necessaria para utilização do env
func ConfigurarEnv(caminhoArquivoEnv string) {
	env.DefinirCaminhoArquivoEnv(caminhoArquivoEnv)
	env.CarregarDotEnv()
}
