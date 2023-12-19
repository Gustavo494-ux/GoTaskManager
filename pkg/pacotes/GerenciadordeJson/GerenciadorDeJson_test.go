package GerenciadordeJson_test

import (
	"GoTaskManager/internal/app/inicializar"
	"GoTaskManager/pkg/pacotes/GerenciadordeJson"
	"GoTaskManager/pkg/pacotes/logger"
	"fmt"
	"os"
	"testing"
)

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	inicializar.InicializarParaTestes()

	exitCode := m.Run()

	if exitCode == 0 {
		logger.Logger().Info("Testes do pacote GerenciadordeJson executados com sucesso!")
	} else {
		logger.Logger().Alerta("Ocorreram erros ao executar os testes do pacote GerenciadordeJson")
	}

	os.Exit(exitCode)
}

func TestInterfaceParaJsonString(t *testing.T) {
	t.Parallel()

	type Pessoa struct {
		Nome  string `json:"nome"`
		Idade int    `json:"idade"`
	}
	p := Pessoa{"João", 30}

	jsonStr, err := GerenciadordeJson.InterfaceParaJsonString(p)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função InterfaceParaJsonString", err)
		t.FailNow()
	}

	expected := `{"nome":"João","idade":30}`
	if jsonStr != expected {
		logger.Logger().Error("Teste "+t.Name()+":	A função InterfaceParaJsonString não retornou o valor esperado",
			err,
			fmt.Sprintf("JSON esperado: %s, JSON retornado: %s", expected, jsonStr),
		)
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func TestJsonStringParaInterface(t *testing.T) {
	t.Parallel()

	jsonStr := `{"nome":"João","idade":30}`
	expected := map[string]interface{}{
		"nome":  "João",
		"idade": float64(30),
	}

	jsonData, err := GerenciadordeJson.JsonStringParaInterface(jsonStr)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função JsonStringParaInterface", err)
		t.FailNow()
	}

	if !compareMaps(jsonData.(map[string]interface{}), expected) {
		logger.Logger().Error("Teste "+t.Name()+":	A função JsonStringParaInterface não retornou o valor esperado",
			err,
			fmt.Sprintf("JSON esperado: %s, JSON retornado: %s", expected, jsonStr),
		)
		t.FailNow()
	}

	logger.Logger().Info("Teste " + t.Name() + ":	Executado com sucesso!")
}

func compareMaps(m1, m2 map[string]interface{}) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}
