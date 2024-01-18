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
func Login(c echo.Context) (err error) {
	var login *models.Login

	if err = json.NewDecoder(c.Request().Body).Decode(&login); err != nil {
		return ResponderErro(c, http.StatusBadRequest, err)
	}

	if err = models.ValidarDados(login); err != nil {
		return ResponderErro(c, http.StatusBadRequest, err)
	}

	token, err := services.Login(login)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao realizar o login", err, login)
		return ResponderErro(c, http.StatusInternalServerError, err)
	}

	return ResponderString(c, http.StatusOK, token)
}
