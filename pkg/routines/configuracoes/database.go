package configuracoes

import (
	"GoTaskManager/pkg/pacotes/database"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

var (
	//BancoProducao: banco de dados destinado a uso efetivo do programa
	BancoProducao *database.ConfiguracaoBancoDeDados

	//BancoTeste: banco de dados destiando a uso durante os testes automatizados
	BancoTeste *database.ConfiguracaoBancoDeDados
)

// ConfigurarNovoBanco: Carrega configura os dados para conexão com o banco
func ConfigurarNovoBanco(host, nomeBanco, usuario, senha, nomeDoDriver, sslmode string, porta int) *database.ConfiguracaoBancoDeDados {
	return database.Novo(
		host,
		nomeBanco,
		usuario,
		senha,
		nomeDoDriver,
		sslmode,
		porta,
	)
}

// ConectarGorm: conecta com o banco de dados utilizando o GORM
func ConectarGorm(c *database.ConfiguracaoBancoDeDados) *gorm.DB {
	return c.ConectarGorm()
}

// ConectarSQLX: conecta com o banco de dados utilizando o SQLX
func ConectarSQLX(c *database.ConfiguracaoBancoDeDados) *sqlx.DB {
	return c.ConectarSQLX()
}

// CriarTabelasGORM: cria as tabelas do banco de dados utilizando o GORM
func CriarTabelasGORM(c *database.ConfiguracaoBancoDeDados, tabelas []interface{}) {
	if c != nil {
		database.CriarTabelasGORM(c.ConectarGorm(), tabelas)
	}
}
