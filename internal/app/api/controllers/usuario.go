package controllers

import (
	"GoTaskManager/internal/app/models"
	"GoTaskManager/internal/app/services"
	"encoding/json"
	"net/http"

	"github.com/Gustavo494-ux/PacotesGolang/logger"
	"github.com/labstack/echo/v4"
)

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(c echo.Context) (err error) {
	var novoUsuario *models.Usuario
	json.NewDecoder(c.Request().Body).Decode(&novoUsuario)
	if err = services.CriarUsuario(novoUsuario); err != nil {

		logger.Logger().Error("Ocorreu um erro criar o usuário", err, novoUsuario)
		return c.JSON(http.StatusBadRequest, err)
	}

	resposta := map[string]interface{}{
		"mensagem": "O usuário foi criado com sucesso!",
	}

	return c.JSON(http.StatusOK, resposta)
}

// BuscarUsuarioPorId encontra um usuário no banco de dados por ID.
func BuscarUsuarioPorId(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Rota em desenvolvimento")
}

// BuscarUsuarioPorEmail encontra um usuário no banco de dados por Email.
func BuscarUsuarioPorEmail(c echo.Context) error {
	email := c.Param("email")
	usuario, err := services.BuscarUsuarioPorEmail(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if usuario == nil {
		return ResponderString(c, http.StatusNotFound, "usuário não encontrado")
	}

	return c.JSON(http.StatusOK, usuario)
}

// BuscarTodosUsuarios recupera todos os usuários salvos no banco de dados.
func BuscarTodosUsuarios(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Rota em desenvolvimento")
}

// AtualizarUsuario atualiza as informações do usuário no banco de dados.
func AtualizarUsuario(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Rota em desenvolvimento")
}

// DeletarUsuario Deleta um usuário do banco de dados.
func DeletarUsuario(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Rota em desenvolvimento")
}
