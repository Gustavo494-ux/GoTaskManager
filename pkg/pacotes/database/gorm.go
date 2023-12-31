package database

import (
	"GoTaskManager/pkg/pacotes/logger"

	"gorm.io/gorm"
)

// ConectarSQLX: Conecta ao banco de dados utilizando o ORM GORM
func (c *ConfiguracaoBancoDeDados) ConectarGorm() (db *gorm.DB) {
	conexaoConfigurada := c.configurarConexaoGORM()

	db, err := gorm.Open(conexaoConfigurada, &gorm.Config{})
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao conectar com o banco de dados utilizando GORM", err, c.StringConexao)

	}
	logger.Logger().Rastreamento("Conexão com o banco de dados estabelecida com sucesso utilizando GORM")
	return
}

// configurarConexaoGORM: configurarConexaoGORM configura a conexão do banco de dados de acordo com o banco de dados passado no NomeDoDriver
func (c *ConfiguracaoBancoDeDados) configurarConexaoGORM() (conexao gorm.Dialector) {
	switch c.NomeDoDriver {
	case "postgres":
		conexao = c.configurarConexaoPostgresGORM()
	}
	logger.Logger().Fatal("Nenhuma configuração de banco de dados encontrada para "+c.NomeDoDriver, nil)
	return
}
