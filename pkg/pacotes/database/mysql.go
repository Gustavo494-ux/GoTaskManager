package database

import "fmt"

// setarStringConexaoMysql: setarStringConexaoMysql configura a string de conexão do mysql
func (c *ConfiguracaoBancoDeDados) setarStringConexaoMysql() {
	c.StringConexao = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.Usuario,
		c.Senha,
		c.Host,
		c.Porta,
		c.NomeBancoDeDados,
	)
}
