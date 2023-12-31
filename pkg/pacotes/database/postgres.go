package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setarStringConexaoPostgres: setarStringConexaoPostgres configura a string de conexão do postgres
func (c *ConfiguracaoBancoDeDados) setarStringConexaoPostgres() {

	c.StringConexao = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%ds sslmode=%s",
		c.Host,
		c.Usuario,
		c.Senha,
		c.NomeBancoDeDados,
		c.Porta,
		c.SSLMode,
	)
}

// configurarConexaoPostgresGORM: configura a conexão do GORM para o banco postgres
func (c *ConfiguracaoBancoDeDados) configurarConexaoPostgresGORM() gorm.Dialector {
	c.setarStringConexaoPostgres()
	return postgres.Open(c.StringConexao)
}
