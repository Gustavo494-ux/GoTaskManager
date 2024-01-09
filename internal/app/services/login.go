package services

import (
	"GoTaskManager/internal/app/models"
	"GoTaskManager/internal/app/repositorio"
	"GoTaskManager/internal/utils"
	"GoTaskManager/pkg/pacotes/authentication"
	"errors"
	"time"
)

// CriarUsuario: cria um novo usuário no banco de dados
func Login(l *models.Login) (token string, err error) {
	if err = models.ValidarDados(l); err != nil {
		return
	}

	usuarioBanco := repositorio.BuscarUsuarioPorEmail(l.Email)
	if usuarioBanco == nil {
		err = errors.New("usuario não encontrado")
		return
	}

	if usuarioBanco.ID == 0 {
		err = errors.New("usuario não encontrado")
		return
	}

	if !utils.CompararHASH(usuarioBanco.Senha, l.Senha) {
		err = errors.New("credenciais incorretas")
		return
	}

	token, err = authentication.NovoToken(true, time.Now().Add(time.Hour*6).Unix()).Criar()
	return
}
