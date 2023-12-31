package database

type ConfiguracaoBancoDeDados struct {
	StringConexao    string
	Host             string
	NomeBancoDeDados string
	Usuario          string
	Senha            string
	NomeDoDriver     string
	SSLMode          string
	Porta            int
}

// Novo: retorna uma nova instância de configuracaoBancoDeDados
func Novo(host, nomeBanco, usuario, senha, nomeDoDriver, sslmode string, porta int) *ConfiguracaoBancoDeDados {
	return &ConfiguracaoBancoDeDados{
		Host:             host,
		NomeBancoDeDados: nomeBanco,
		Usuario:          usuario,
		Senha:            senha,
		NomeDoDriver:     nomeDoDriver,
		SSLMode:          sslmode,
		Porta:            porta,
	}
}
