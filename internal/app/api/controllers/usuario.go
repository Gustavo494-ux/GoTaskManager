package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CriarUsuario insere um usuário no banco de dados.
func CriarUsuario(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Rota em desenvolvimento")
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
