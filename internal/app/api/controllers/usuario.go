package controllers

import (
	"net/http"

	"github.com/Gustavo494-ux/PacotesGolang/logger"
	"github.com/labstack/echo/v4"

	"GoTaskManager/internal/app/models"
	"GoTaskManager/internal/app/services"
	"GoTaskManager/internal/utils"
)

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(c echo.Context) (err error) {
	var novoUsuario models.Usuario

	if err = utils.ExtrairBodyEmStruct(c.Request(), &novoUsuario); err != nil {
		return c.String(http.StatusBadRequest, "Ocorreu um erro ao desserializar o corpo da requisição")
	}

	if err = services.CriarUsuario(&novoUsuario); err != nil {
		return c.String(http.StatusInternalServerError, "Ocorreu um erro ao criar o usuário")
	}

	return c.String(http.StatusCreated, "O usuário foi criado com sucesso!")
}

// BuscarUsuarioPorId encontra um usuário no banco de dados por ID.
func BuscarUsuarioPorId(c echo.Context) error {
	idUsuario := uint(RetornarParametroInteiro(c, "usuarioId"))

	usuario, err := services.BuscarUsuarioPorId(idUsuario)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, usuario)
}

// BuscarUsuarioPorEmail encontra um usuário no banco de dados por Email.
func BuscarUsuarioPorEmail(c echo.Context) error {
	usuario, err := services.BuscarUsuarioPorEmail(c.Param("email"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, usuario)
}

// BuscarTodosUsuarios recupera todos os usuários salvos no banco de dados.
func BuscarTodosUsuarios(c echo.Context) error {
	usuarios, err := services.BuscarTodosUsuarios()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, usuarios)
}

// AtualizarUsuario atualiza as informações do usuário no banco de dados.
func AtualizarUsuario(c echo.Context) error {
	var usuario models.Usuario
	idUsuario := uint(RetornarParametroInteiro(c, "usuarioId"))

	if err := utils.ExtrairBodyEmStruct(c.Request(), &usuario); err != nil {
		logger.Logger().Error("Ocorreu um erro ao desserializar o corpo da requisição", err, usuario)
		return c.String(
			http.StatusBadRequest,
			"ocorreu um erro ao desserializar o corpo da requisição",
		)
	}

	if err := services.AtualizarUsuario(&usuario, idUsuario); err != nil {
		return c.String(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	return c.NoContent(http.StatusOK)
}

// DeletarUsuario Deleta um usuário do banco de dados.
func DeletarUsuario(c echo.Context) error {
	if err := services.DeletarUsuario(uint(RetornarParametroInteiro(c, "usuarioId"))); err != nil {
		return c.String(
			http.StatusBadRequest,
			err.Error(),
		)
	}
	return c.NoContent(http.StatusOK)
}
