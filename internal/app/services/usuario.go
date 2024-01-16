package services

import (
	"GoTaskManager/internal/app/models"
	"GoTaskManager/internal/app/repositorio"
	"GoTaskManager/internal/utils"
	"GoTaskManager/pkg/pacotes/GerenciadordeJson"
	"GoTaskManager/pkg/pacotes/logger"
	"GoTaskManager/pkg/routines/configuracoes"
	"encoding/json"
	"errors"
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

// DeletarUsuario: deleta um usuário do banco de dados
func DeletarUsuario(id uint) (err error) {
	return configuracoes.BancoPrincipalGORM.Delete(&models.Usuario{}, id).Error
}

func BuscarUsuarioPorId(id uint) (usuario models.Usuario, err error) {
	err = configuracoes.BancoPrincipalGORM.First(&usuario, id).Error
	return
}

// TratarUsuarioParaResposta: trata o usuário para responder a solicitação de forma adequada
func TratarUsuarioParaResposta(usuarioInput *models.Usuario) {
	jsonByte, err := GerenciadordeJson.IgnorarCamposPelaTag(*usuarioInput, "serializar", "false")
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao remover os campos com a tag serializar contendo o valor 'false' do struct", err, usuarioInput)
		return
	}
	usuarioTemporario := models.Usuario{}

	err = json.Unmarshal(jsonByte, &usuarioTemporario)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao desserializar o json para o struct de usuário", err, jsonByte, usuarioInput)
		return
	}

	*usuarioInput = usuarioTemporario
}
