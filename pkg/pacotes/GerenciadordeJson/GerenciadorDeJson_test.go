package GerenciadordeJson_test

import (
	"GoTaskManager/pkg/pacotes/GerenciadordeJson"
	"GoTaskManager/pkg/pacotes/logger"
	"GoTaskManager/pkg/routines/inicializarpkg"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// TestMain:Função executada antes das demais
func TestMain(m *testing.M) {
	inicializarpkg.InicializarLogger()
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

func TestIgnorarCamposPelaTag(t *testing.T) {
	type MeuStruct struct {
		Campo1 string `minhaTag:"valor1"`
		Campo2 string `minhaTag:"valor2"`
		Campo3 string `minhaTag:"valor1"`
	}

	meuObjeto := MeuStruct{
		Campo1: "valor1",
		Campo2: "valor2",
		Campo3: "valor1",
	}

	j, err := GerenciadordeJson.IgnorarCamposPelaTag(meuObjeto, "minhaTag", "valor1")
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao executar a função IgnorarCamposPelaTag", err)
		t.FailNow()
	}

	if !json.Valid(j) {
		logger.Logger().Error("Teste "+t.Name()+":	A função IgnorarCamposPelaTag não retornou um JSON válido", nil, string(j))
		t.FailNow()
	}

	var m map[string]interface{}
	err = json.Unmarshal(j, &m)
	if err != nil {
		logger.Logger().Error("Teste "+t.Name()+":	Ocorreu um erro ao deserializar o JSON", err, string(j))
		t.FailNow()
	}

	_, existe := m["Campo1"]
	if existe {
		logger.Logger().Error("Teste "+t.Name()+":	A função IgnorarCamposPelaTag não ignorou os campos que atendem à condição", err)
		t.FailNow()
	}

	_, existe = m["Campo2"]
	if !existe {
		logger.Logger().Error("Teste "+t.Name()+":	A função IgnorarCamposPelaTag não incluiu os campos que não atendem à condição", err)
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
