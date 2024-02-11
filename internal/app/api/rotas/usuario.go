package rotas

import (
	"github.com/labstack/echo/v4"

	"GoTaskManager/internal/app/api/controllers"
	"GoTaskManager/internal/app/api/middlewares"
	middlewaresUsuario "GoTaskManager/internal/app/api/middlewares/usuario"
)

// RotasUsuario: Criação das rotas de usuário
func RotasUsuario(e *echo.Echo) {
	e.POST("/usuario", controllers.CriarUsuario, middlewaresUsuario.ValidarUsuarioInput)

	userGroup := e.Group("/usuario")
	userGroup.Use(middlewares.Authenticate)

	userGroup.GET("", controllers.BuscarTodosUsuarios)
	userGroup.GET("/id/:usuarioId", controllers.BuscarUsuarioPorId)
	userGroup.GET("/:email", controllers.BuscarUsuarioPorEmail)

	userGroup.PUT("/:usuarioId", controllers.AtualizarUsuario, middlewaresUsuario.ValidarUsuarioInput)
	userGroup.DELETE("/:usuarioId", controllers.DeletarUsuario)
}
