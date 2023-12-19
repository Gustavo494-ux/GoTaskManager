package authentication

import (
	"GoTaskManager/internal/app/inicializar"
	"GoTaskManager/pkg/pacotes/logger"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	inicializar.InicializarParaTestes()
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
	tokenBuilder := NovoToken(true, time.Now().Add(time.Hour*6).Unix())
	if len(tokenBuilder.Claims) == 0 {
		logger.Logger().Alerta("Teste: " + t.Name() + ":	a função NovoToken não retornou um TokenBuilder")
		t.FailNow()
	}
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestAdicionarParametro(t *testing.T) {
	t.Parallel()

	tokenBuilder := NovoToken(true, time.Now().Add(time.Hour*6).Unix())
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
		NovoToken(true, time.Now().Add(time.Hour*6).Unix()).
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
		NovoToken(true, time.Now().Add(time.Hour*6).Unix()).
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

	err = ValidarToken(*req)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função ValidarToken", err)
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestExtrairInformacao(t *testing.T) {
	t.Parallel()
	token, _ := NovoToken(true, time.Now().Add(time.Hour*6).Unix()).
		AdicionarParametro("usuarioId", 1).
		Criar()

	usuarioIdString, err := ExtrairInformacao(token, "usuarioId")
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função ExtrairInformacao", err)
		t.FailNow()
	}

	usuarioId, err := strconv.Atoi(usuarioIdString)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao converter o usuarioId de tipo string em int", err)
		t.FailNow()
	}

	if !reflect.DeepEqual(usuarioId, 1) {
		logger.Logger().Alerta("Teste " + t.Name() + ":	A função extrairInformacao não retornou o valor esperado")
		t.FailNow()
	}
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestExtrairTodasInformacoes(t *testing.T) {
	t.Parallel()
	tokenBuilder := NovoToken(true, time.Now().Add(time.Hour*6).Unix())
	var parametrosEntrada = map[string]string{
		"usuarioId":  "1",
		"permissao1": "true",
		"permissao2": "false",
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

	parametrosSaida, err := ExtrairTodasInformacoes(token)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função ExtrairTodasInformacoes", err)
		t.FailNow()
	}

	for chave := range parametrosEntrada {
		if parametrosEntrada[chave] != parametrosSaida[chave] {
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

	token := ExtrairToken(*req)

	if token != "abc123" {
		logger.Logger().Error("Teste "+t.Name()+":	A função ExtrairToken não retornou o valor esperado", nil)
		t.FailNow()
	}
	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}
