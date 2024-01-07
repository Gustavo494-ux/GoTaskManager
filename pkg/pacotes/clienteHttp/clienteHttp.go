package clientehttp

import (
	"GoTaskManager/pkg/pacotes/logger"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	contentType = "application/json"
)

type request struct {
	resposta   *http.Response
	err        error
	body       []byte
	statusCode int
}

// POST: Realiza uma requisição POST
func POST(URL string, body any) (requisicao *request) {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro converter o body em json", err, body)
		return
	}

	response, err := http.Post(URL, contentType, bytes.NewBuffer(bodyJson))
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao realizar uma requisição POST", err, URL, body)
		return
	}

	requisicao = &request{
		resposta: response,
		err:      err,
	}
	return
}

// GetBody: retorna o corpo da requisição como slice de bytes
func (requisicao *request) GetBody() (body []byte, err error) {
	body, err = io.ReadAll(requisicao.resposta.Body)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ler a resposta de uma requisição", err, requisicao.resposta)
		return
	}
	requisicao.body = body
	return
}

// GetStatusCode: retorna o status code da requisição
func (requisicao *request) GetStatusCode() (statusCode int) {
	requisicao.statusCode = requisicao.resposta.StatusCode
	return requisicao.statusCode
}
