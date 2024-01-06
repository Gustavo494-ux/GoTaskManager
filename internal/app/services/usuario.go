package services

import (
	"GoTaskManager/internal/app/models"
	"GoTaskManager/internal/app/repositorio"
	"GoTaskManager/pkg/pacotes/encryption"
	"GoTaskManager/pkg/pacotes/logger"
	"errors"
)

// CriarUsuario: cria um novo usuário no banco de dados
func CriarUsuario(u *models.Usuario) (err error) {
	if u == nil {
		return errors.New("usuário não informado")
	}

	if err = models.ValidarDados(u); err != nil {
		return
	}

	if err = DefinirEmailHash(u); err != nil {
		return
	}

	if err = repositorio.CriarUsuario(u); err != nil {
		return
	}

	return
}

// DefinirEmailHash: atribui ao campo emailHash o valor do campo email após passar pelo hash
func DefinirEmailHash(u *models.Usuario) (err error) {
	u.Email_Hash, err = encryption.GerarSHA512(u.Email)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao encriptar o email do usuario", err, u)
		return
	}
	return
}
