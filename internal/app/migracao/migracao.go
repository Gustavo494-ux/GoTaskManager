package migracao

import (
	"GoTaskManager/internal/app/models"

	"github.com/Gustavo494-ux/PacotesGolang/configuracoes"
)

// CriacaoAutomaticaTabelas: criação automatica das tabelas no GORM
func CriacaoAutomaticaTabelas() {
	var tabelas []interface{}

	tabelas = append(tabelas, models.Usuario{})

	configuracoes.CriarTabelasGORM(configuracoes.BancoPrincipalGORM, tabelas)
}
