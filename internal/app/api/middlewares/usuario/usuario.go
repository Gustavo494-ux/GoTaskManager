package middlewaresUsuario

import (
	"net/http"

	_ "github.com/Gustavo494-ux/PacotesGolang/tipo"
	"github.com/labstack/echo/v4"

	"GoTaskManager/internal/app/api/middlewares"
	"GoTaskManager/internal/app/models"
	"GoTaskManager/internal/app/services"
	"GoTaskManager/internal/utils"
)

// ValidarUsuarioInput: valida os dados do usuário antes de entrar no controller
func ValidarUsuarioInput(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		usuario := models.Usuario{}
		utils.ExtrairBodyEmStruct(c.Request(), &usuario)

		if err := utils.ValidarBodyModel(usuario); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if err := services.VerificarSeUsuarioExiste(usuario); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return middlewares.ProximaFuncao(c, next)
	}
}
