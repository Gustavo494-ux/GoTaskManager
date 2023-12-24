package rotas

import (
	"GoTaskManager/internal/app/api/controllers"

	"github.com/labstack/echo/v4"
)

// RotasUsuario: Criação das rotas de usuário
func RotasUsuario(e *echo.Echo) {
	userGroup := e.Group("/usuario")

	userGroup.POST("", controllers.CriarUsuario)
	userGroup.GET("", controllers.BuscarTodosUsuarios)
	userGroup.GET("/:usuarioId", controllers.BuscarUsuarioPorId)
	userGroup.GET("/:usuarioEmail", controllers.BuscarUsuarioPorEmail)
	userGroup.PUT("/:usuarioId", controllers.AtualizarUsuario)
	userGroup.DELETE("/:usuarioId", controllers.DeletarUsuario)
}
