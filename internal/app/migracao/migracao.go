package migracao

import (
	"GoTaskManager/internal/app/models"
	"GoTaskManager/pkg/routines/configuracoes"
)

// CriacaoAutomaticaTabelas: criação automatica das tabelas no GORM
func CriacaoAutomaticaTabelas() {
	var tabelas []interface{}

	tabelas = append(tabelas, models.Usuario{})

	configuracoes.CriarTabelasGORM(configuracoes.BancoProducao, tabelas)
	configuracoes.CriarTabelasGORM(configuracoes.BancoTeste, tabelas)
}
