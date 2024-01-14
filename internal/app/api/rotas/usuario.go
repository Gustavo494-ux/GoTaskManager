package rotas

import (
	"GoTaskManager/internal/app/api/controllers"
	"GoTaskManager/internal/app/api/middlewares"

	"github.com/labstack/echo/v4"
)

// RotasUsuario: Criação das rotas de usuário
func RotasUsuario(e *echo.Echo) {
	e.POST("/usuario", controllers.CriarUsuario)

	userGroup := e.Group("/usuario")
	userGroup.Use(middlewares.Authenticate)

	userGroup.GET("", controllers.BuscarTodosUsuarios)
	userGroup.GET("/:usuarioId", controllers.BuscarUsuarioPorId)
	userGroup.GET("/:email", controllers.BuscarUsuarioPorEmail)
	userGroup.PUT("/:usuarioId", controllers.AtualizarUsuario)
	userGroup.DELETE("/:usuarioId", controllers.DeletarUsuario)
}
