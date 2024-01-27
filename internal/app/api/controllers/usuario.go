package controllers

import (
	"GoTaskManager/internal/app/models"
	"GoTaskManager/internal/app/services"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Gustavo494-ux/PacotesGolang/logger"
	"github.com/labstack/echo/v4"
)

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(c echo.Context) (err error) {
	var novoUsuario *models.Usuario
	json.NewDecoder(c.Request().Body).Decode(&novoUsuario)
	if err = services.CriarUsuario(novoUsuario); err != nil {
		logger.Logger().Error("Ocorreu um erro criar o usuário", err, novoUsuario)
		return ResponderErro(c, http.StatusBadRequest, err)
	}
	return ResponderString(c, http.StatusOK, "O usuário foi criado com sucesso!")
}

// BuscarUsuarioPorId encontra um usuário no banco de dados por ID.
func BuscarUsuarioPorId(c echo.Context) error {
	usuarioId, err := strconv.Atoi(c.Param("usuarioId"))
	if err != nil {
		return ResponderErro(c, http.StatusBadRequest, err)
	}

	usuario, err := services.BuscarUsuarioPorId(uint(usuarioId))
	if err != nil {
		return ResponderErro(c, http.StatusInternalServerError, err)
	}

	return ResponderUsuario(c, http.StatusOK, "", usuario)
}

// BuscarUsuarioPorEmail encontra um usuário no banco de dados por Email.
func BuscarUsuarioPorEmail(c echo.Context) error {
	email := c.Param("email")
	usuario, err := services.BuscarUsuarioPorEmail(email)
	if err != nil {
		return ResponderErro(c, http.StatusBadRequest, err)
	}

	return ResponderUsuario(c, http.StatusOK, "", usuario)
}

// BuscarTodosUsuarios recupera todos os usuários salvos no banco de dados.
func BuscarTodosUsuarios(c echo.Context) error {
	usuarios, err := services.BuscarTodosUsuarios()
	if err != nil {
		return ResponderErro(c, http.StatusInternalServerError, err)
	}
	return ResponderUsuario(c, http.StatusOK, "", usuarios...)
}

// AtualizarUsuario atualiza as informações do usuário no banco de dados.
func AtualizarUsuario(c echo.Context) error {
	usuarioId, err := strconv.Atoi(c.Param("usuarioId"))
	if err != nil {
		return ResponderErro(c, http.StatusBadRequest, err)
	}

	var usuario models.Usuario

	if err := json.NewDecoder(c.Request().Body).Decode(&usuario); err != nil {
		logger.Logger().Error("Ocorreu um erro ao desserializar o corpo da requisição", err, usuario)
		return ResponderErro(c,
			http.StatusBadRequest,
			errors.New("Ocorreu um erro ao desserializar o corpo da requisição"),
		)
	}

	if err := services.AtualizarUsuario(&usuario, uint(usuarioId)); err != nil {
		return ResponderErro(c,
			http.StatusBadRequest,
			err,
		)
	}

	return ResponderUsuario(c, http.StatusOK, "usuário atualizado")
}

// DeletarUsuario Deleta um usuário do banco de dados.
func DeletarUsuario(c echo.Context) error {
	usuarioId, err := strconv.Atoi(c.Param("usuarioId"))
	if err != nil {
		return ResponderErro(c, http.StatusBadRequest, err)
	}

	if err := services.DeletarUsuario(uint(usuarioId)); err != nil {
		return ResponderErro(c,
			http.StatusBadRequest,
			err,
		)
	}
	return ResponderString(c, http.StatusOK, "usuário deletado")
}

func ResponderUsuario(c echo.Context, statusCodeSucesso int, mensagemSucesso string, usuarios ...*models.Usuario) (err error) {
	existeUsuario := false
	for _, usuario := range usuarios {

		if usuario == nil {
			continue
		}

		if usuario.ID == 0 {
			continue
		}
		existeUsuario = true

	}

	if !existeUsuario {
		return ResponderString(c, http.StatusNotFound, "nenhum usuário não encontrado")
	}

	if mensagemSucesso != "" {
		return ResponderString(c, statusCodeSucesso, mensagemSucesso)
	} else {
		return c.JSON(statusCodeSucesso, usuarios)
	}
}
