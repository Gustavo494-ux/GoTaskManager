package database

import (
	"GoTaskManager/pkg/pacotes/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ConectarSQLX: Conecta ao banco de dados utilizando o pacote SQLX
func (c *ConfiguracaoBancoDeDados) ConectarSQLX() (db *sqlx.DB, err error) {
	c.configurarStringConexaoSQLX()

	db, err = sqlx.Open(c.NomeDoDriver, c.StringConexao)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao conectar com o banco de dados utilizando SQLX", err, c.StringConexao)
	}

	logger.Logger().Rastreamento("Conexão com o banco de dados estabelecida com sucesso utilizando SQLX")
	return
}

// configurarStringConexaoSQLX: configurarStringSQLX configura a string de conexão de acordo com o driver do banco
func (c *ConfiguracaoBancoDeDados) configurarStringConexaoSQLX() {
	switch c.NomeDoDriver {
	case "mysql":
		c.setarStringConexaoMysql()
	}
}
