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
func Login(c echo.Context) (err error) {
	var login *models.Login

	if err = utils.ExtrairBodyEmStruct(c.Request(), login); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err = models.ValidarDados(login); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	token, err := services.Login(login)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao realizar o login", err, login)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, token)
}
