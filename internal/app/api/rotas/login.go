package rotas

import (
	"GoTaskManager/internal/app/api/controllers"

	"github.com/labstack/echo/v4"
)

// RotasLogin: Criação das rotas de login
func RotasLogin(e *echo.Echo) {
	e.POST("/login", controllers.Login)
}
