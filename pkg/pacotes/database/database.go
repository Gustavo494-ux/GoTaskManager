package database

import (
	"GoTaskManager/pkg/pacotes/logger"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	DriverName    string
	StringConexao string
)

func ConfigurarConexao(nomeDriver, stringConexao string) {
	DriverName = nomeDriver
	StringConexao = stringConexao
}

// Conectar: Conecta ao banco de dados
func Conectar() (db *sqlx.DB, err error) {
	db, err = sqlx.Open(DriverName, StringConexao)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao conectar com o banco de dados", err, StringConexao)
	}

	logger.Logger().Rastreamento("Conexão com o banco de dados estabelecida com sucesso")
	return
}
