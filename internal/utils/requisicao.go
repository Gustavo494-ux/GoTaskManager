package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Gustavo494-ux/PacotesGolang/logger"

	"GoTaskManager/internal/app/models"
)

// RetornarCorpoRequisicao: ler e retorna o corpo da requisição porém consumir o mesmo.
func RetornarCorpoRequisicao(Request *http.Request) (corpo []byte, err error) {
	corpo, err = io.ReadAll(Request.Body)
	if err != nil {
		return
	}
	defer Request.Body.Close()

	Request.Body = io.NopCloser(bytes.NewBuffer(corpo))
	return
}

// ExtrairBodyEmStruct: extrai os dados do body para um objeto. o parametro dados deve ser um ponteiro
// o parametro dados deve ser exatamente do tipo de dados esperado na requisição. Caso seja esperado uma lista o tipo deve ser um slice.
func ExtrairBodyEmStruct(Request *http.Request, dados any) {
	corpo, err := RetornarCorpoRequisicao(Request)
	if err != nil {
		logger.Logger().Error("Ocorreu um erro ao extrair o corpo da requisição", err)
	}

	if err := json.Unmarshal(corpo, dados); err != nil {
		logger.Logger().Error("Ocorreu um erro ao desserializar o body da requisição", err, corpo)
	}
}

// ValidarBodyModel: valida o body da requisição como um model
func ValidarBodyModel(a ...any) (err error) {
	for _, o := range a {
		err = models.ValidarDados(o)
		if err != nil {
			logger.Logger().Info("o body fornecido é invalido", o, err)
			return fmt.Errorf("o body fornecido é invalido.	%s", err.Error())
		}
	}
	return
}
