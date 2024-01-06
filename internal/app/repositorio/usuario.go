package repositorio

import (
	"GoTaskManager/internal/app/models"
	"GoTaskManager/pkg/pacotes/logger"
	"GoTaskManager/pkg/routines/configuracoes"
)

// CriarUsuario: cria um novo usuário no banco de dados
func CriarUsuario(u *models.Usuario) (err error) {
	if err = configuracoes.BancoPrincipalGORM.Create(u).Error; err != nil {
		logger.Logger().Error("Ocorreu um erro ao criar o usuário no banco", err, u)
	}
	return
}
