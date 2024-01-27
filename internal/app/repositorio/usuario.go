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

// BuscarUsuarioPorCPF: busca um usuário no banco de dados pelo seu CPF
func BuscarUsuarioPorCPF(cpf string) (usuario *models.Usuario) {
	configuracoes.BancoPrincipalGORM.Where("cpf =?", cpf).First(&usuario)
	return
}

// BuscarTodosUsuarios: busca todos os usuários do banco de dados
func BuscarTodosUsuarios() (usuarios []*models.Usuario) {
	configuracoes.BancoPrincipalGORM.Find(&usuarios)
	return
}

// BuscarUsuarioPorId: busca o usuário utilizando o ID
func BuscarUsuarioPorId(id uint) (usuarios *models.Usuario) {
	configuracoes.BancoPrincipalGORM.First(&usuarios, id)
	return
}

// AtualizarUsuario: atualiza dos dados do usuário no banco de dados
func AtualizarUsuario(usuario *models.Usuario) (err error) {
	if err = configuracoes.BancoPrincipalGORM.Save(usuario).Error; err != nil {
		logger.Logger().Error("Ocorreu um erro ao realizar a atualização do usuário", err, usuario)
	}
	return
}

// DeletarUsuario: deleta o usuário do banco de dados
func DeletarUsuario(usuario *models.Usuario) (err error) {
	if err = configuracoes.BancoPrincipalGORM.Delete(usuario).Error; err != nil {
		logger.Logger().Error("Ocorreu um erro ao deletar o usuário", err, usuario)
	}
	return
}
