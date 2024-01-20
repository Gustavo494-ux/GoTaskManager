package services

import (
	"GoTaskManager/internal/app/models"
	"GoTaskManager/internal/app/repositorio"
	"GoTaskManager/internal/utils"
	"encoding/json"
	"errors"

	"github.com/Gustavo494-ux/PacotesGolang/GerenciadordeJson"
	"github.com/Gustavo494-ux/PacotesGolang/logger"
)

// CriarUsuario: cria um novo usuário no banco de dados
func CriarUsuario(u *models.Usuario) (err error) {
	if err = models.ValidarDados(u); err != nil {
		return
	}

	u.Email_Hash = utils.GerarHash(u.Email)
	u.Senha = utils.GerarHash(u.Senha)

	if err = repositorio.CriarUsuario(u); err != nil {
		return
	}
	return
}

// BuscarUsuarioPorEmail: busca um usuário no banco de dados pelo seu email
func BuscarUsuarioPorEmail(email string) (usuario *models.Usuario, err error) {
	if email == "" {
		err = errors.New("email não informado")
	}
	usuario = repositorio.BuscarUsuarioPorEmail(email)
	TratarUsuarioParaResposta(usuario)
	return
}

// BuscarTodosUsuarios: busca todos os usuários do banco de dados
func BuscarTodosUsuarios() (usuarios []*models.Usuario, err error) {
	usuarios = repositorio.BuscarTodosUsuarios()
	TratarUsuarioParaResposta(usuarios...)
	return
}

// BuscarUsuarioPorId: busca o usuário utilizando o ID
func BuscarUsuarioPorId(id uint) (usuario *models.Usuario, err error) {
	usuario = repositorio.BuscarUsuarioPorId(id)
	TratarUsuarioParaResposta(usuario)
	return
}

// TratarUsuarioParaResposta: trata o usuário para responder a solicitação de forma adequada
func TratarUsuarioParaResposta(usuariosInput ...*models.Usuario) {
	for _, usuarioInput := range usuariosInput {
		if usuarioInput == nil {
			return
		}

		if usuarioInput.ID == 0 {
			usuarioInput = nil
			return
		}

		jsonByte, err := GerenciadordeJson.IgnorarCamposPelaTag(*usuarioInput, "serializar", "false")
		if err != nil {
			logger.Logger().Error("Ocorreu um erro ao remover os campos com a tag serializar contendo o valor 'false' do struct", err, usuarioInput)
			return
		}
		*usuarioInput = models.Usuario{Model: usuarioInput.Model}
		err = json.Unmarshal(jsonByte, usuarioInput)
		if err != nil {
			logger.Logger().Error("Ocorreu um erro ao desserializar o json para o struct de usuário", err, jsonByte, usuarioInput)
			return
		}
	}
	return
}
