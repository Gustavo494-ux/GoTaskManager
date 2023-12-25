package configuracoes

import "GoTaskManager/pkg/pacotes/env"

//ConfigurarEnv: realiza a configuração necessaria para utilização do env
func ConfigurarEnv(caminhoArquivoEnv string) {
	env.DefinirCaminhoArquivoEnv(caminhoArquivoEnv)
	env.CarregarDotEnv()
}
