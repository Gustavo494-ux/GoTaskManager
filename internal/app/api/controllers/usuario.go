package controllers

import (
	"GoTaskManager/pkg/pacotes/authentication"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(c echo.Context) error {
	token, _ := authentication.NovoToken(true, time.Now().Add(time.Second*10).Unix()).AdicionarParametro("Usuario", "Gustavo").Criar()
	return c.JSON(http.StatusAccepted, token)

	// return c.JSON(http.StatusNotFound, "Rota em desenvolvimento")
}

// BuscarUsuarioPorId encontra um usuário no banco de dados por ID.
func BuscarUsuarioPorId(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Rota em desenvolvimento")
}

// BuscarUsuarioPorEmail encontra um usuário no banco de dados por Email.
func BuscarUsuarioPorEmail(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Rota em desenvolvimento")
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
