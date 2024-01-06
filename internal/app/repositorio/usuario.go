package repositorio

import (
	"GoTaskManager/internal/app/models"
	"GoTaskManager/pkg/routines/configuracoes"
)

// CriarUsuario: cria um novo usuário no banco de dados
func CriarUsuario(c *models.Usuario) (err error) {
	db := configuracoes.BancoProducao.ConectarGorm()
	return db.Create(c).Error
}
