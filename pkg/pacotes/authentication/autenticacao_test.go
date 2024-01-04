package authentication_test

import (
	"GoTaskManager/pkg/pacotes/authentication"
	"GoTaskManager/pkg/pacotes/logger"
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
)

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	authentication.DefinirSecretKey("secrectKey")

	exitCode := m.Run()
	if exitCode == 0 {
		logger.Logger().Info("Testes do pacote authentication executados com sucesso!")
	} else {
		logger.Logger().Alerta("Ocorreram erros ao executar os testes do pacote authentication")
	}
	os.Exit(exitCode)
}

func TestNovoToken(t *testing.T) {
	t.Parallel()
	tokenBuilder := authentication.NovoToken(true, time.Now().Add(time.Hour*6).Unix())
	if len(tokenBuilder.Claims) == 0 {
		logger.Logger().Alerta("Teste: " + t.Name() + ":	a função NovoToken não retornou um TokenBuilder")
		t.FailNow()
	}
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestAdicionarParametro(t *testing.T) {
	t.Parallel()

	tokenBuilder := authentication.NovoToken(true, time.Now().Add(time.Hour*6).Unix())
	tokenBuilder.AdicionarParametro("usuarioId", 1)

	if _, ok := tokenBuilder.Claims["usuarioId"]; !ok {
		logger.Logger().Alerta("Teste " + t.Name() + ":	a função AdicionarParametro não adicionou o parâmetro ao token")
		t.FailNow()
	}

	logger.Logger().Info("Teste: " + t.Name() + "	Executado com sucesso!")
}

func TestCriar(t *testing.T) {
	t.Parallel()
	token, err :=
		authentication.NovoToken(true, time.Now().Add(time.Hour*6).Unix()).
			AdicionarParametro("usuarioId", 1).
			Criar()

	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao criar o token", err)
		t.FailNow()
	}

	if reflect.DeepEqual(token, "") {
		logger.Logger().Alerta("Teste " + t.Name() + ":	A função criar não retornou um token")
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestValidarToken(t *testing.T) {
	t.Parallel()
	token, err :=
		authentication.NovoToken(true, time.Now().Add(time.Hour*6).Unix()).
			AdicionarParametro("usuarioId", 1).
			Criar()

	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função Criar", err)
		t.FailNow()
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro criar um request", err)
		t.FailNow()
	}

	req.Header.Set("Authorization", "Bearer "+token)

	err = authentication.ValidarToken(*req)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função ValidarToken", err)
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestExtrairInformacao(t *testing.T) {
	t.Parallel()
	token, _ := authentication.NovoToken(true, time.Now().Add(time.Hour*6).Unix()).
		AdicionarParametro("usuarioId", 1).
		Criar()

	usuarioIdInterface, err := authentication.ExtrairInformacao(token, "usuarioId")
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função ExtrairInformacao", err)
		t.FailNow()
	}

	usuarioIdJsonSaida, err := json.Marshal(usuarioIdInterface)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao converter usuarioIdJsonSaida para json", err)
		t.FailNow()
	}

	usuarioIdJsonEntrada, err := json.Marshal(1)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao converter usuarioIdJsonEntrada para json", err)
		t.FailNow()
	}

	if string(usuarioIdJsonEntrada) != string(usuarioIdJsonSaida) {
		logger.Logger().Alerta("Teste " + t.Name() + ":	A função extrairInformacao não retornou o valor esperado")
		t.FailNow()
	}
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestExtrairTodasInformacoes(t *testing.T) {
	t.Parallel()
	tokenBuilder := authentication.NovoToken(true, time.Now().Add(time.Hour*6).Unix())
	parametrosEntrada := map[string]interface{}{
		"usuarioId":  1,
		"permissao1": true,
		"permissao2": false,
		"tipo":       "teste",
	}

	for chave, valor := range parametrosEntrada {
		tokenBuilder.AdicionarParametro(chave, valor)
	}

	token, err := tokenBuilder.Criar()
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função Criar()", err)
		t.FailNow()
	}

	parametrosSaida, err := authentication.ExtrairTodasInformacoes(token)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função ExtrairTodasInformacoes", err)
		t.FailNow()
	}

	var entrada, saida []byte
	for chave := range parametrosEntrada {
		entrada, err = json.Marshal(parametrosEntrada[chave])
		if err != nil {
			logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao converter parametrosEntrada[chave] para json, na chave "+chave, err)
			t.FailNow()
		}

		saida, err = json.Marshal(parametrosSaida[chave])
		if err != nil {
			logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao converter parametrosSaida[chave] para json, na chave "+chave, err)
			t.FailNow()
		}

		if string(entrada) != string(saida) {
			logger.Logger().Alerta("Teste " + t.Name() + ":	A função ExtrairTodasInformacoes não retornou o valor esperado")
			t.FailNow()
		}
	}
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestExtrairToken(t *testing.T) {
	t.Parallel()

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer abc123")

	token := authentication.ExtrairToken(*req)

	if token != "abc123" {
		logger.Logger().Error("Teste "+t.Name()+":	A função ExtrairToken não retornou o valor esperado", nil)
		t.FailNow()
	}
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}
