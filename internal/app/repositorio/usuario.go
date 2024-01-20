package repositorio

import (
	"GoTaskManager/internal/app/models"

	"github.com/Gustavo494-ux/PacotesGolang/configuracoes"
	"github.com/Gustavo494-ux/PacotesGolang/logger"
)

// CriarUsuario: cria um novo usuário no banco de dados
func CriarUsuario(u *models.Usuario) (err error) {
	if err = configuracoes.BancoPrincipalGORM.Create(u).Error; err != nil {
		logger.Logger().Error("Ocorreu um erro ao criar o usuário no banco", err, u)
	}
	return
}

// BuscarUsuarioPorEmail: busca um usuário no banco de dados pelo seu email
func BuscarUsuarioPorEmail(email string) (usuario *models.Usuario) {
	configuracoes.BancoPrincipalGORM.Where("email =?", email).First(&usuario)
	return
}

// BuscarTodosUsuarios: busca todos os usuários do banco de dados
func BuscarTodosUsuarios() (usuarios []*models.Usuario) {
	configuracoes.BancoPrincipalGORM.Find(&usuarios)
	return
}
