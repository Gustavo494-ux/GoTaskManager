package configuracoes

import (
	"GoTaskManager/pkg/pacotes/database"
	"GoTaskManager/pkg/pacotes/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

var (
	BancoPrincipalGORM *gorm.DB

	//ConfiguracaoBancoProducao: banco de dados destinado a uso efetivo do programa
	ConfiguracaoBancoProducao *database.ConfiguracaoBancoDeDados

	//BancoTeste: banco de dados destiando a uso durante os testes automatizados
	ConfiguracaoBancoTeste *database.ConfiguracaoBancoDeDados
)

// ConfigurarNovoBanco: Carrega configura os dados para conexão com o banco
func ConfigurarNovoBanco(host, nomeBanco, usuario, senha, nomeDoDriver, sslmode string, porta string) *database.ConfiguracaoBancoDeDados {
	portaInt, err := strconv.Atoi(porta)
	if err != nil {
		logger.Logger().Fatal("Ocorreu um erro ao converter a PORTA_DATABASE_TESTE para string", err)
	}
	return database.Novo(
		host,
		nomeBanco,
		usuario,
		senha,
		nomeDoDriver,
		sslmode,
		portaInt,
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
func CriarTabelasGORM(db *gorm.DB, tabelas []interface{}) {
	if db != nil {
		database.CriarTabelasGORM(db, tabelas)
	}
}
